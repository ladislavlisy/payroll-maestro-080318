package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	server := NewServer()

	log.Println("Server started: http://localhost:" + port)
	server.Run(":" + port)
}
