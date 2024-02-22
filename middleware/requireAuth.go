package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lordofthemind/ginJwtAuthPostgres/initializers"
	"github.com/lordofthemind/ginJwtAuthPostgres/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("RequireAuth middleware")

	// Get the cookie off request

	tokenString, err := c.Cookie("Authorisation")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Abort()
		return
	}

	// Check if the cookie is valid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//Check the expiration date

		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token subject

		var user models.User
		initializers.DB.First(&user, "id = ?", claims["subject"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to request context
		c.Set("user", user)
		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Abort()
		return
	}
}
