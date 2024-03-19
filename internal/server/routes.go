package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// r.Get("/api/healthz", s.healthHandler)

	apiV1Router := chi.NewRouter()
	r.Mount("/api/v1", apiV1Router)
	apiV1Router.Mount("/users", s.userRouter())
	apiV1Router.Get("/", s.HelloWorldHandler)

	// Auth routes
	apiV1Router.Get("/auth/{provider}/callback", s.getAuthCallbackHandler)
	apiV1Router.Get("/logout/{provider}", s.LogoutHandler)
	apiV1Router.Get("/auth/{provider}", s.authHandler)

	return r
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	err := s.writeJSON(w, http.StatusOK, helloWorldResponse{"Hello World"})
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
}

// func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
// 	jsonResp, _ := json.Marshal(s.db.Health())
// 	_, _ = w.Write(jsonResp)
// }
