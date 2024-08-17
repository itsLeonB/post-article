package main

import (
	"log"
	"post-api/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env: %e", err)
	}

	a := app.Init()
	a.Serve()
}
