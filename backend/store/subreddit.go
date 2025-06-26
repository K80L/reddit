package store

import "gorm.io/gorm"

type Subreddit struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(32);column:name;unique;index"`

	Posts []Post `gorm:"foreignKey:SubredditID"`
	// Users []User `gorm:"foreignKey:ID"` // many-to-many relationship w/ Users
}

type SubredditStore struct {
	db *gorm.DB
}

func NewSubredditStore(db *gorm.DB) (*SubredditStore, error) {
	return &SubredditStore{db: db}, nil
}

func (s *SubredditStore) CreateSubreddit(subreddit *Subreddit) error {
	result := s.db.Create(subreddit)

	return result.Error
}

func (s *SubredditStore) GetSubreddit(id int) (*Subreddit, error) {
	var subreddit Subreddit
	result := s.db.Preload("Posts").Where("id = ?", id).First(&subreddit)

	return &subreddit, result.Error
}
