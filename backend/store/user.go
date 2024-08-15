package store

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(255);column:username;unique"`
	Password string `gorm:"not null;type:varchar(255);column:password"`
	Email    string `gorm:"not null;type:varchar(255);column:email;unique"`

	Likes    []Like    `gorm:"foreignKey:UserID"`
	Dislikes []Dislike `gorm:"foreignKey:UserID"`
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
	result := db.Where("username = ?", username).First(&user)

	return &user, result.Error
}

func GetUserById(id int) (*User, error) {
	var user User

	result := db.Preload("Likes").Preload("Dislikes").Where("id = ?", id).First(&user)

	return &user, result.Error
}
