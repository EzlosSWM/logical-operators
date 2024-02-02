package main

import (
	"log"
	"os"
	"strconv"

	"engineering/internal/server"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := server.NewServer()

	log.Println("Listening on port: ", port)
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
