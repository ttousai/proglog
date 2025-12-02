package main

import (
	"log"

	"github.com/ttousai/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPSever(":8080")
	log.Fatal(srv.ListenAndServe())
}
