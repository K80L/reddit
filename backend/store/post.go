package store

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string `gorm:"not null;type:varchar(100);column:title"`
	Content string `gorm:"not null;type:text;column:content"`

	Likes   []Like    `gorm:"foreignKey:PostID"`
	Disikes []Dislike `gorm:"foreignKey:PostID"`

	// Posts can have one parent post and many children posts
	ParentID *int   `gorm:"column:parent_id" json:"parent_id"`
	Parent   *Post  `gorm:"foreignKey:ParentID;references:ID"`
	Children []Post `gorm:"foreignKey:ParentID;references:ID"`

	// Posts belong to a subreddit
	SubredditID int       `gorm:"not null;column:subreddit_id" json:"subreddit_id"`
	Subreddit   Subreddit `gorm:"foreignKey:SubredditID"`

	// Posts belong to a user
	UserID int  `gorm:"not null;column:user_id" json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
}

func GetPosts() ([]Post, error) {
	var posts []Post
	result := db.Preload("User").Preload("Subreddit").Preload("Children").Preload("Parent").Find(&posts)

	return posts, result.Error
}

func GetPostByID(postID int) (Post, error) {
	var post Post
	result := db.Preload("User").Where("id = ?", postID).First(&post)

	return post, result.Error
}

func CreatePost(post *Post) error {
	result := db.Create(post)

	return result.Error
}

func DeletePost(id int) error {
	// soft delete
	result := db.Delete(Post{}, id) // Is this how to do a Delete??

	return result.Error
}
