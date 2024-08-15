package store

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID int  `gorm:"not null;column:user_id"`
	User   User `gorm:"foreignKey:UserID;references:UserID"`
	PostID int  `gorm:"not null;colunn:post_id"`
	Post   Post `gorm:"foreignKey:PostID;references:PostID"`
}

type Dislike struct {
	gorm.Model
	UserID int  `gorm:"not null;column:user_id"`
	User   User `gorm:"foreignKey:UserID;references:UserID"`
	PostID int  `gorm:"not null;colunn:post_id"`
	Post   Post `gorm:"foreignKey:PostID;references:PostID"`
}

func LikePost(postId, userId int) error {
	post := Post{}

	result := db.Where("id = ?", postId).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Like{UserID: userId, PostID: postId})
	return result.Error
}

func DislikePost(postId, userId int) error {
	post := Post{}

	result := db.Where("id = ?", postId).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Dislike{UserID: userId, PostID: postId})
	return result.Error
}
