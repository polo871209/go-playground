package main

import (
	"fmt"
	"net/http"

	"github.com/polo871209/chi-playground/internal/auth"
	"github.com/polo871209/chi-playground/internal/database"
)

type authHanlder func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHanlder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
