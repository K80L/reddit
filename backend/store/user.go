package store

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(255);column:username;unique"`
	Password string `gorm:"not null;type:varchar(255);column:password"`
	Email    string `gorm:"not null;type:varchar(255);column:email;unique"`

	Posts    []Post    `gorm:"foreignKey:UserID"`
	Likes    []Like    `gorm:"foreignKey:UserID"`
	Dislikes []Dislike `gorm:"foreignKey:UserID"`

	// Subreddits []Subreddit `gorm:"foreignKey:ID"` // many-to-many relationship w/ Subreddits
}

func AddUser(user *User) error {
	encryptedPassword, err := EncryptPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = encryptedPassword
	db.Create(user)

	return nil
}

func GetUser(username string) (*User, error) {
	var user User
	result := db.Preload("Posts").Preload("Likes").Preload("Dislikes").Where("username = ?", username).First(&user)

	return &user, result.Error
}

func GetUserById(id int) (*User, error) {
	var user User

	result := db.Preload("Posts").Preload("Likes").Preload("Dislikes").Where("id = ?", id).First(&user)

	return &user, result.Error
}

func (u *User) HasLiked(postID int) bool {
	for _, like := range u.Likes {
		fmt.Println("like.PostID", like.PostID)
		if like.PostID == postID {
			return true
		}
	}

	return false
}

func (u *User) HasDisliked(postID int) bool {
	for _, dislike := range u.Dislikes {
		if dislike.PostID == postID {
			return true
		}
	}

	return false
}
