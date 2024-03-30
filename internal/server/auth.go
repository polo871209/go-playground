package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

type User struct {
	Provider          string
	Email             string
	Name              string
	NickName          string
	AvatarURL         string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

// getAuthCallbackHandler godoc
//
//	@Summary		Processes the authentication callback
//	@Description	This endpoint is the callback URL for Google authentication.
//
// It attempts to complete the user authentication process with Google. On success, it redirects the user; on failure, it redirects to an error page.
// Note: Currently, only Google is supported as a third-party provider.
//
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			provider	path		string			true			"The name of the third-party provider"	Enums(google)
//	@Success		200			{object}	User			@description:	Represents								the		authenticated	user's	information	from	Google.
//	@Failure		403			{object}	errorResponse	@description:	An										error	response		object	detailing	why		authentication	failed.
//	@Router			/auth/{provider}/callback [get]
func (s *Server) getAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	_, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		s.errorJSON(w, err, http.StatusForbidden)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}

// LogoutHandler godoc
//
//	@Summary		Logs out the user
//	@Description	Logs out the user from the current session by clearing authentication cookies or tokens, then redirects to the home page.
//
// Note: Currently, only Google is supported for logout functionality.
//
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			provider	path		string	true	"The name of the third-party provider"	Enums(google)
//	@Success		302			{string}	string	"User is redirected to the home page."
//	@Router			/auth/{provider}/logout [get]
func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

// authHandler godoc
//
//	@Summary		Login with a third-party provider
//	@Description	Initiates authentication with a specified third-party provider and returns user information upon success.
//
// Note: Currently, only Google is supported as a third-party provider for authentication.
//
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			provider	path		string	true			"The name of the third-party provider"	Enums(google)
//	@Success		200			{object}	User	@description:	Represents								the	authenticated	user's	information.
//	@Router			/auth/{provider} [get]
func (s *Server) authHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	if _, err := gothic.CompleteUserAuth(w, r); err == nil {

		http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

// AuthMiddleware checks if the user is authenticated.
// If not, it sends an HTTP 403 Forbidden response.
func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated using Goth
		if _, err := gothic.CompleteUserAuth(w, r); err != nil {
			// User is not authenticated
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// User is authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
