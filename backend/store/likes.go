package store

type Like struct {
	BaseModel
	UserID int64 `gorm:"not null;column:user_id"`
	User   User  `gorm:"foreignKey:UserID;references:ID"`
	PostID int64 `gorm:"not null;colunn:post_id"`
	Post   Post  `gorm:"foreignKey:PostID;references:ID"`
}

type Dislike struct {
	BaseModel
	UserID int64 `gorm:"not null;column:user_id"`
	User   User  `gorm:"foreignKey:UserID;references:ID"`
	PostID int64 `gorm:"not null;colunn:post_id"`
	Post   Post  `gorm:"foreignKey:PostID;references:ID"`
}

func LikePost(postId, userId int64) error {
	post := Post{}

	result := db.Where("id = ?", postId).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Like{UserID: userId, PostID: postId})
	return result.Error
}

func DislikePost(postId, userId int64) error {
	post := Post{}

	result := db.Where("id = ?", postId).First(&post)

	if result.Error != nil {
		return result.Error
	}

	db.Create(&Dislike{UserID: userId, PostID: postId})
	return result.Error
}
