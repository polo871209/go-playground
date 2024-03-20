package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type gothUser = goth.User

// getAuthCallbackHandler godoc
// @Summary Processes the authentication callback
// @Description This endpoint is the callback URL for Google authentication.
// It attempts to complete the user authentication process with Google. On success, it redirects the user; on failure, it redirects to an error page.
// Note: Currently, only Google is supported as a third-party provider.
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "The name of the third-party provider" Enums(google)
// @Success 200 {object} goth.User @description: Represents the authenticated user's information from Google.
// @Failure 403 {object} errorResponse @description: An error response object detailing why authentication failed.
// @Router /api/auth/{provider}/callback [get]
func (s *Server) getAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		s.errorJSON(w, err, http.StatusForbidden)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	log.Println(user)

	s.writeJSON(w, http.StatusOK, user)
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

// LogoutHandler godoc
// @Summary Logs out the user
// @Description Logs out the user from the current session by clearing authentication cookies or tokens, then redirects to the home page.
// Note: Currently, only Google is supported for logout functionality.
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "The name of the third-party provider" Enums(google)
// @Success 302 {string} string "User is redirected to the home page."
// @Router /api/auth/{provider}/logout [get]
func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

// authHandler godoc
// @Summary Login with a third-party provider
// @Description Initiates authentication with a specified third-party provider and returns user information upon success.
// Note: Currently, only Google is supported as a third-party provider for authentication.
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "The name of the third-party provider" Enums(google)
// @Success 200 {object} gothUser @description: Represents the authenticated user's information.
// @Router /api/auth/{provider} [get]
func (s *Server) authHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		s.writeJSON(w, http.StatusOK, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
