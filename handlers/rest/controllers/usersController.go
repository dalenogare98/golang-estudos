package controllers

import (
	"net/http"
	"os"
	"time"

	"crud-go/initializers"
	"crud-go/middleware"
	"crud-go/models"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c echo.Context) error {
	// Get the email/pass off req body
	var body models.User

	err := c.Bind(&body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to read body",
		})
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to hash password",
		})
	}
	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to create user",
		})
	}
	// Respond
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func Login(c echo.Context) error {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to read body",
		})
	}

	// Look up the user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid email or password",
		})
	}

	// Comparte sent in pass with
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": ("Failed to create token"),
		})
	}

	// Send it back
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{})

}

func Validate(c echo.Context) error {
	return middleware.RequireAuth(c)
}
