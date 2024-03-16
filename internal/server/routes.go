package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	// Auth routes
	r.Get("/auth/{provider}/callback", s.getAuthCallbackHandler)
	r.Get("/logout/{provider}", s.LogoutHandler)
	r.Get("/auth/{provider}", s.authHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

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
