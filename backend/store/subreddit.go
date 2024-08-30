package store

import "gorm.io/gorm"

type Subreddit struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(32);column:name;unique;index"`

	Posts []Post `gorm:"foreignKey:SubredditID"`
	// Users []User `gorm:"foreignKey:ID"` // many-to-many relationship w/ Users
}

func CreateSubreddit(subreddit *Subreddit) error {
	result := db.Create(subreddit)

	return result.Error
}

func GetSubreddit(id int) (*Subreddit, error) {
	var subreddit Subreddit
	result := db.Preload("Posts").Where("id = ?", id).First(&subreddit)

	return &subreddit, result.Error
}
