package store

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

func CreateJWT(user *User) string {
	// Get the secret from the environment
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Create the claims with user data
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func (s *UserStore) Authenticate(username, password string) (*User, error) {
	user, err := s.GetUser(username)

	if !ComparePassword(password, user.Password) {
		return nil, fmt.Errorf("Invalid password")
	}

	return user, err
}

func ValidateJWT(c *gin.Context) (string, error) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("No token cookie found")
	}

	tokenString := cookie.Value
	if tokenString == "" {
		return "", fmt.Errorf("Unauthorized")
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("Unauthorized")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}

	return username, nil
}

// Encrypts a password using bcrypt
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Compares a password with a hash
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
