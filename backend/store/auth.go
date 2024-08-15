package store

import (
	"fmt"
	"log"
	"os"
	"time"

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

func Authenticate(username, password string) (*User, error) {
	user, err := GetUser(username)

	if !ComparePassword(password, user.Password) {
		return nil, fmt.Errorf("Invalid password")
	}

	return user, err
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
