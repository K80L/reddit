package main

import (
	"fmt"
	"log"

	"github.com/K80L/reddit/backend/server"
	"github.com/K80L/reddit/backend/store"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err := store.Init()

	if err != nil {
		fmt.Println(err)
	}

	server.Init()
}
