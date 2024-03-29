package auth

import (
	"errors"
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/request_bodies"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func register_new_user(db *gorm.DB, email string, username string, password string, confirm string) error {
	/* Check if email is already taken */
	var user1 = &notebook_db.User{}

	db.Where("email = ?", email).First(&user1)

	if email == user1.Email {
		return errors.New("email already taken")
	}

	/* Check if username is already taken */
	var user2 = &notebook_db.User{}

	db.Where("username = ?", username).First(&user2)

	if username == user2.Username {
		return errors.New("username already taken")
	}

	// Create user account
	notebook_db.CreateNewUser(db, email, username, password)

	return nil
}

func Register(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Register */
		var body request_bodies.RegisterRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Register new user
		err = register_new_user(db, body.Email, body.Username, body.Password, body.ConfirmPwd)

		/* Check if user registration is successful */
		if err == nil {
			// Send success message
			c.JSON(200, gin.H{
				"registered":  true,
				"msg_success": "Registered Successfully",
			})
		} else {
			// Send error message
			c.JSON(200, gin.H{
				"registered": false,
				"msg_error":  err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
