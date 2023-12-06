package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/polo871209/chi-playground/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// if string(params.FeedID) == "" {
	// 	respondWithError(w, http.StatusBadRequest, "Name cannot be empty")
	// 	return
	// }

	feedFollowChan := make(chan database.FeedFollow)
	errChan := make(chan error)

	go func() {
		defer close(feedFollowChan)
		defer close(errChan)

		feedFellow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC().Add(time.Hour * 8),
			UpdatedAt: time.Now().UTC().Add(time.Hour * 8),
			FeedID:    params.FeedID,
			UserID:    user.ID,
		})

		if err != nil {
			errChan <- err
			return
		}
		feedFollowChan <- feedFellow
	}()

	select {
	case feedFollow := <-feedFollowChan:
		respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
	case err := <-errChan:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error follow feed: %v", err))
	}
}

func (apiCfg *apiConfig) handlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFellows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFellows))
}
