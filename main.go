package main

import (
	"log"

	"github.com/iverian/advent-of-code-2023-go/day4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("unable to load dotenv file: %s", err)
	}

	if err := day4.Main(); err != nil {
		log.Fatalf("%s", err)
	}
}
