package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func Protect(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token cookie found"})
		c.Abort()
		return
	}

	tokenString := cookie.Value

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims := jwt.MapClaims{}
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := os.Getenv("JWT_SECRET")

		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	username := claims["username"].(string)
	user, err := store.GetUser(username)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	c.Set("user", user)
}

func CustomErrors(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		fmt.Println(c.Errors)
		for _, err := range c.Errors {
			fmt.Println(err.Type)
			switch err.Type {
			case gin.ErrorTypePublic:
				if !c.Writer.Written() {
					c.AbortWithStatusJSON(c.Writer.Status(), gin.H{"error": err.Error()})
				}

			case gin.ErrorTypeBind:
				errMap := make(map[string]string)
				if errs, ok := err.Err.(validator.ValidationErrors); ok {
					for _, fieldErr := range []validator.FieldError(errs) {
						errMap[fieldErr.Field()] = customValidationError(fieldErr)
					}
				}

				status := http.StatusBadRequest
				// preserve current status
				if c.Writer.Status() != http.StatusOK {
					status = c.Writer.Status()
				}
				c.AbortWithStatusJSON(status, gin.H{"error": errMap})
			default:
				log.Error().Err(err.Err).Msg("Unknown error")
			}
		}

		if !c.Writer.Written() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}
}

func customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", err.Field())
	case "min":
		return fmt.Sprintf("%s must be longer than or equal %s characters.", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters.", err.Field(), err.Param())
	default:
		return err.Error()
	}
}
