package main

import (
	"log"
	"net/http"
	sw "github.com/chuhongwei/BlogServer/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
