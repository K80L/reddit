package store

type UserStorage interface {
	Authenticate(username, password string) (*User, error)
	AddUser(user *User) error
	GetUser(username string) (*User, error)
	GetUserById(id int) (*User, error)
}

type SubredditStorage interface {
	CreateSubreddit(subreddit *Subreddit) error
	GetSubreddit(id int) (*Subreddit, error)
}

type PostStorage interface {
	GetPosts() ([]Post, error)
	GetPostByID(postID int) (Post, error)
	CreatePost(post *Post) error
	DeletePost(id int) error
}

type LikeStorage interface {
	LikePost(postID, userID int) error
	DislikePost(postID, userID int) error
	Undo(postID, userID int) error
	HasLiked(postID int) bool
	HasDisliked(postID int) bool
}
