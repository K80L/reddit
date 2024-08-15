package store

type Subreddit struct {
	BaseModel
	Name string `gorm:"not null;type:varchar(32);column:name;unique;index"`

	Posts []Post `gorm:"foreignKey:SubredditID;references:ID"`
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
