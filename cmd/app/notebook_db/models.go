package notebook_db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email		string
	Username	string
	Password	string
	Note		[]Note
}

type Note struct {
	gorm.Model
	Title       string
	Description string
	UserID		uint
}

/* User functions */

func CreateNewUser(db *gorm.DB, email string, username string, password string) {
	db.Create(&User{Email: email, Username: username, Password: password})
}

/* Note functions */

func GetNoteEntries(db *gorm.DB) []Note {
	/* Get all the entries of the notes table */
	var notes []Note // products list
	db.Find(&notes)  // find entries of notes table

	return notes
}

func CreateNewNoteEntry(db *gorm.DB, title string, description string) {
	/* Create new note entry */
	db.Create(&Note{Title: title, Description: description})
}

func GetNoteEntry(db *gorm.DB, id string) *Note {
	/* Get the entry by id */
	var note = &Note{}
	db.First(&note, id)
	return note
}

func UpdateNoteEntry(db *gorm.DB, id string, title string, description string) {
	/* Update the entry's title and description by id */
	var note = &Note{}

	// Find the first record that matches the id
	db.First(&note, id)

	// Update Title and Description text
	db.Model(&note).Updates(Note{Title: title, Description: description})
}
