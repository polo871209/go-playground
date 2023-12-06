package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/polo871209/chi-playground/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name cannot be empty")
		return
	}

	feedChan := make(chan database.Feed)
	errChan := make(chan error)

	go func() {
		defer close(feedChan)
		defer close(errChan)

		feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC().Add(time.Hour * 8),
			UpdatedAt: time.Now().UTC().Add(time.Hour * 8),
			Name:      params.Name,
			Url:       params.URL,
			UserID:    user.ID,
		})

		if err != nil {
			errChan <- err
			return
		}
		feedChan <- feed
	}()

	select {
	case feed := <-feedChan:
		respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
	case err := <-errChan:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error create user: %v", err))
	}
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
