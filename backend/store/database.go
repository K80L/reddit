package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db = db.Debug()

	//error := db.Migrator().DropTable(&Post{}, &Subreddit{}, &User{})

	// if err != nil {
	// 	fmt.Println(error)
	// }

	if err != nil {
		panic("failed to connect database")
	}

	models := []interface{}{&User{}, &Subreddit{}, &Post{}, &Like{}, &Dislike{}}

	for _, model := range models {
		fmt.Printf("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMigrating model: %+v\n", model)
		if err := db.AutoMigrate(model); err != nil {
			panic("failed to migrate model")
		}
	}

	return db, nil
}

func GetConnection() *gorm.DB {
	return db
}
