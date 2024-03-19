package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (s *Server) getAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	log.Println(user.AvatarURL)
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *Server) authHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Println("already login")
		log.Println(gothUser.AvatarURL)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
