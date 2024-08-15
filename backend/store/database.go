package store

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type BaseModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func Init() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	//error := db.Migrator().DropTable(&Post{}, &Subreddit{}, &User{})

	// if err != nil {
	// 	fmt.Println(error)
	// }
	if err != nil {
		panic("failed to connect database")
	}

	models := []interface{}{&User{}, &Subreddit{}, &Like{}, &Dislike{}, &Post{}}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			panic("failed to migrate model")
		}
	}

	return db, nil
}

func GetConnection() *gorm.DB {
	return db
}
