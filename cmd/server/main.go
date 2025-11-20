package main

import (
	"log"

	"github.com/travisjefferey/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPSever(":8080")
	log.Fatal(srv.ListenAndServe())
}
