package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/polo871209/go-playground/internal/database"
)

// createUser godoc
//	@Summary	create user by name
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	201	{object}	database.User
//	@Failure	400	{object}	errorResponse
//	@Router		/users [post]
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

	w.WriteHeader(http.StatusCreated)
	s.writeJSON(w, http.StatusCreated, user)
}
