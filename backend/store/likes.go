package store

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID int  `gorm:"not null;column:user_id"`
	User   User `gorm:"foreignKey:UserID"`
	PostID int  `gorm:"not null;column:post_id"`
	Post   Post `gorm:"foreignKey:PostID"`
}

type Dislike struct {
	gorm.Model
	UserID int  `gorm:"not null;column:user_id"`
	User   User `gorm:"foreignKey:UserID"`
	PostID int  `gorm:"not null;column:post_id"`
	Post   Post `gorm:"foreignKey:PostID"`
}

func LikePost(postID, userID int) error {
	post := Post{}

	result := db.Where("id = ?", postID).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Like{UserID: userID, PostID: postID})
	return result.Error
}

func DislikePost(postID, userID int) error {
	post := Post{}

	result := db.Where("id = ?", postID).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Dislike{UserID: userID, PostID: postID})
	return result.Error
}

func UndoLikePost(postID, userID int) error {
	like := Like{}

	result := db.Where("post_id", postID).Where("user_id", userID).First(&like)

	if result.Error != nil {
		return result.Error
	}

	db.Delete(&like)

	return result.Error
}

func UndoDislikePost(postID, userID int) error {
	dislike := Dislike{}

	result := db.Where("post_id", postID).Where("user_id", userID).First(&dislike)

	if result.Error != nil {
		return result.Error
	}

	db.Delete(&dislike)

	return result.Error
}
