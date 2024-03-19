package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/polo871209/go-playground/internal/database"
)

func (s *Server) userRouter() *chi.Mux {
	userRouter := chi.NewRouter()
	userRouter.Post("/", s.createUser) // Use the server's createUser method
	return userRouter
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}

	err := s.readJSON(w, r, &params)
	if err != nil {
		s.errorJSON(w, err)
		return
	}

	user, err := s.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})

	if err != nil {
		s.errorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.writeJSON(w, http.StatusCreated, user)
}
