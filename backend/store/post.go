package store

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string `gorm:"not null;type:varchar(100);column:title"`
	Content string `gorm:"not null;type:text;column:content"`

	Likes []Like `gorm:"foreignKey:PostID"`

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

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) (*PostStore, error) {
	return &PostStore{db: db}, nil
}

func (s *PostStore) GetPosts() ([]Post, error) {
	var posts []Post
	result := s.db.Preload("User").Preload("Subreddit").Preload("Children").Preload("Parent").Find(&posts)

	return posts, result.Error
}

func (s *PostStore) GetPostByID(postID int) (Post, error) {
	var post Post
	result := s.db.Preload("User").Where("id = ?", postID).First(&post)

	return post, result.Error
}

func (s *PostStore) CreatePost(post *Post) error {
	result := s.db.Create(post)

	return result.Error
}

func (s *PostStore) DeletePost(id int) error {
	// soft delete
	result := s.db.Delete(Post{}, id) // Is this how to do a Delete??

	return result.Error
}
