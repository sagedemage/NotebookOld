package router_setup

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"go-web-app-experiment/cmd/app/notebook_db"

	"go-web-app-experiment/cmd/app/notes"

	"go-web-app-experiment/cmd/app/template_loader"

	"go-web-app-experiment/cmd/app/page_renderer"

	"go-web-app-experiment/cmd/app/auth"
)

func InitializeRouter() *gin.Engine {
	// create the router
	router := gin.Default()

	// Open database
	db := notebook_db.InitDB("database/notebook.db")

	// html renderer
	router.HTMLRender = template_loader.LoadTemplates("cmd/app/templates")

	// session
	store := cookie.NewStore([]byte("secret"))
  	router.Use(sessions.Sessions("mysession", store))

	// do not trust all proxies for security reasons
	router.SetTrustedProxies(nil)

	// Load static files (for css, and etc)
	router.Static("/static", "cmd/app/static")

	/* Get Requests */
	// Render the home page at the root of the website
	router.GET("/", page_renderer.RenderWebPage("home.tmpl", "Home")) // home.tmpl does not exist

	// Render the about page at the route "/about"
	router.GET("/about", page_renderer.RenderWebPage("about.tmpl", "About"))

	// Render the new registration page at route "/register"
	router.GET("/register", page_renderer.RenderWebPage("register.tmpl", "Register"))

	// Render the login page at route "/login"
	router.GET("/login", page_renderer.RenderWebPage("login.tmpl", "Login"))

	/* Post Requests */
	// Register the user
	router.POST("/register", auth.Register(db))

	// Login the user
	router.POST("/login", auth.Login(db))

	/* Auhtorization Required */
	auth_routes := router.Group("/").Use(auth.AuthRequired)

	/* Get Requets */
	// Render the view table page at route "/table"
	auth_routes.GET("/view-notes", notes.ViewNotes(db))

	// Render the new entry page at route "/new-entry"
	auth_routes.GET("/add-new-note", page_renderer.RenderWebPage("new-note.tmpl", "New Note"))

	// Render the new entry page at route "/new-entry"
	auth_routes.GET("/edit-note", notes.EditNoteForm(db))

	// Logout the user
	auth_routes.GET("/logout", auth.Logout)

	/* Post Requests */
	// Render the view table page at route "/table"
	auth_routes.POST("/view-notes", notes.DeleteOrEditNote(db))

	// Get Form data from POST request
	auth_routes.POST("/add-new-note", notes.AddNewNote(db))

	// Get Form data from POST request
	auth_routes.POST("/edit-note", notes.EditNote(db))

	return router
}