package main

import (
	"fmt"
	"log"

	"github.com/polo871209/go-playground/internal/auth"
	"github.com/polo871209/go-playground/internal/server"
)

func main() {
	auth.NewAuth()
	server := server.NewServer()
	log.Printf("starting server on port %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
