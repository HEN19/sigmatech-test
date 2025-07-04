package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("API-")

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(user model.UserModel) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Id:       user.ID.Int64,
		Username: user.Username.String,
		Name:     user.FirstName.String + " " + user.LastName.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

// AuthMiddleware is the middleware to validate JWT tokens.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		// Validate the token
		_, err := ValidateToken(tokenString)
		if err != nil {
			// Respond with Unauthorized status using the ResponseOut function
			out.ResponseOut(c, nil, false, http.StatusUnauthorized, "Unauthorized")
			// Stop further processing
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

func DecodeToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
