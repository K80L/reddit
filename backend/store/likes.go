package store

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID int    `gorm:"not null;column:user_id"`
	User   User   `gorm:"foreignKey:UserID"`
	PostID int    `gorm:"not null;column:post_id"`
	Post   Post   `gorm:"foreignKey:PostID"`
	Type   string `gorm:"not null;type:varchar(10);column:type"` // "like" or "dislike"
}

type LikeStore struct {
	db *gorm.DB
}

func (s *LikeStore) LikePost(postID, userID int) error {
	post := Post{}

	result := s.db.Where("id = ?", postID).First(&post)

	if result.Error != nil {
		return result.Error
	}

	s.db.Create(&Like{UserID: userID, PostID: postID, Type: "like"})
	return result.Error
}

func (s *LikeStore) DislikePost(postID, userID int) error {
	post := Post{}

	result := s.db.Where("id = ?", postID).First(&post)

	if result.Error != nil {
		return result.Error
	}

	s.db.Create(&Like{UserID: userID, PostID: postID, Type: "dislike"})
	return result.Error
}

func (s *LikeStore) Undo(postID, userID int) error {
	like := Like{}

	result := s.db.Where("post_id", postID).Where("user_id", userID).First(&like)

	if result.Error != nil {
		return result.Error
	}

	s.db.Delete(&like)

	return result.Error
}

// func (u *UserStore) HasLiked(postID int) bool {
// 	for _, like := range u.Likes {
// 		fmt.Println("like.PostID", like.PostID)
// 		if like.PostID == postID {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (u *UserStore) HasDisliked(postID int) bool {
// 	for _, dislike := range u.Dislikes {
// 		if dislike.PostID == postID {
// 			return true
// 		}
// 	}

// 	return false
// }
