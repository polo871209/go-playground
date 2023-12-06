package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/polo871209/chi-playground/internal/database"
)

func (apiCfg *apiConfig) handlerCreteUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, 400, "Name cannot be empty")
		return
	}

	userChan := make(chan database.User)
	errChan := make(chan error)

	go func() {
		defer close(userChan)
		defer close(errChan)

		user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC().Add(time.Hour * 8),
			UpdatedAt: time.Now().UTC().Add(time.Hour * 8),
			Name:      params.Name,
		})

		if err != nil {
			errChan <- err
			return
		}
		userChan <- user
	}()

	select {
	case user := <-userChan:
		respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
	case err := <-errChan:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error create user: %v", err))
	}
}
