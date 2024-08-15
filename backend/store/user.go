package store

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int    `gorm:"not null;column:user_id;unique;autoIncrement;primaryKey"`
	Username string `gorm:"not null;type:varchar(255);column:username;unique"`
	Password string `gorm:"not null;type:varchar(255);column:password"`
	Email    string `gorm:"not null;type:varchar(255);column:email;unique"`
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
