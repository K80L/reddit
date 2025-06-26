package store

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(255);column:username;unique"`
	Password string `gorm:"not null;type:varchar(255);column:password"`
	Email    string `gorm:"not null;type:varchar(255);column:email;unique"`

	Posts []Post `gorm:"foreignKey:UserID"`
	Likes []Like `gorm:"foreignKey:UserID"`

	// Subreddits []Subreddit `gorm:"foreignKey:ID"` // many-to-many relationship w/ Subreddits
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) (*UserStore, error) {
	return &UserStore{db: db}, nil
}

func (s *UserStore) AddUser(user *User) error {
	encryptedPassword, err := EncryptPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = encryptedPassword
	s.db.Create(user)

	return nil
}

func (s *UserStore) GetUser(username string) (*User, error) {
	var user User
	result := s.db.Preload("Posts").Preload("Likes").Where("username = ?", username).First(&user)

	return &user, result.Error
}

func (s *UserStore) GetUserById(id int) (*User, error) {
	var user User

	result := s.db.Preload("Posts").Preload("Likes").Preload("Dislikes").Where("id = ?", id).First(&user)

	return &user, result.Error
}
