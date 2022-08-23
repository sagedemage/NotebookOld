package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/models"
)

/* Get Requests */

func RenderWebPage(template_page string, page_title string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//var template_page string = "home/index.tmpl"
		//var page_title string = "Home Page"

		c.HTML(http.StatusOK, template_page, gin.H{
			"title": page_title,
			"page_title": page_title,
		})
	}
	return gin.HandlerFunc(fn)
}

func ViewNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */

		// Set title name for the page
		var title_name = "Notes"

		// entries of the notes database
		notes := models.GetNoteEntries(db)

		// Pass the list of notes to the web page
		c.HTML(http.StatusOK, "view-notes.tmpl", gin.H{
			"title": title_name,
			"page_title": title_name,
			"note_list":  notes,
		})
	}
	return gin.HandlerFunc(fn)
}

func EditNoteForm(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Set title name for the page
		var title_name = "Edit Note"

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// get entry note values
		note := models.GetNoteEntry(db, id)

		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "notes/edit-note.tmpl", gin.H{
			"title":			title_name,
			"page_title":       title_name,
			"note_title":       note.Title,
			"note_description": note.Description,
		})

	}
	return gin.HandlerFunc(fn)
}


