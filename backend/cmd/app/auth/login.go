package auth

import (
	"errors"
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/request_bodies"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func is_user_valid(db *gorm.DB, username string, password string) (uint, error) {
	/* Check if the User is Valid */
	var user = &notebook_db.User{}

	// Get entry with the specified email or username
	db.Where("email = ? OR username = ?", username, username).First(&user)

	if username == user.Email || username == user.Username {
		/* Check if the email or username exists */
		// compare the password to the password hash
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			/* Check if the password is incorrect */
			return user.ID, errors.New("incorrect username or password")
		}
	} else if username != user.Email || username != user.Username {
		/* Check if the email or username does not exists */
		return user.ID, errors.New("incorrect username or password")
	}

	return user.ID, nil
}

func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Login */
		var body request_bodies.LoginRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Is User Valid
		user_id, err := is_user_valid(db, body.Username, body.Password)

		/* Check if user registration is successful */
		if err == nil {
			token, err := GenerateToken(user_id)

			if err != nil {
				var err = errors.New("failed to generate token")
				println(err)
			}

			c.JSON(200, gin.H{
				"auth":  true,
				"token": token,
			})
		} else {
			// json message
			c.JSON(200, gin.H{
				"auth":      false,
				"msg_error": err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
