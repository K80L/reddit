package store

import "gorm.io/gorm"

type Subreddit struct {
	gorm.Model
	SubredditID int    `gorm:"not null;column:subreddit_id;unique;autoIncrement;"`
	Name        string `gorm:"not null;type:varchar(32);column:name;unique;index"`

	Posts []Post `gorm:"foreignKey:SubredditRefer"`
}

func CreateSubreddit(subreddit *Subreddit) error {
	result := db.Create(subreddit)

	return result.Error
}

func GetSubreddit(name string) (*Subreddit, error) {
	var subreddit Subreddit
	result := db.Preload("Posts").Where("name = ?", name).First(&subreddit)

	return &subreddit, result.Error
}
