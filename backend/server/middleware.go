package server

import (
	"fmt"
	"net/http"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func Protect(userStore *store.UserStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := store.ValidateJWT(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		user, err := userStore.GetUser(username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// set user in context for later use
		c.Set("user", user)
		c.Next()
	}
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
					return
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
				return
			default:
				log.Error().Err(err.Err).Msg("Unknown error")
			}
		}

		if !c.Writer.Written() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
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
