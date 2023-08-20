package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"crud-go/initializers"
	"crud-go/models"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c echo.Context) error {
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "No token found",
		})
		
	}

	// Decode/Validate it
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// c.AbortWithStatus(http.StatusUnauthorized)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Token expired",
			})
			
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		// Attach to req
		c.Set("user", user)

		
	} 
	return c.JSON(http.StatusOK, echo.Map{
		"token": token.Raw,
	})

	// Find the user with token sub
}
