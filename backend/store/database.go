package store

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	var err error
	dbConnectionString := os.Getenv("DATABASE_URL")
	if dbConnectionString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db = db.Debug()

	//error := db.Migrator().DropTable(&Post{}, &Subreddit{}, &User{})

	// if err != nil {
	// 	fmt.Println(error)
	// }

	if err != nil {
		panic("failed to connect database")
	}

	models := []any{&User{}, &Subreddit{}, &Post{}, &Like{}}

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
