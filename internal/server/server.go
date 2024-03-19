package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/polo871209/go-playground/internal/database"
)

type Server struct {
	port int

	db *database.Queries
}

func db() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	return conn
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: port,

		db: database.New(db()),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
