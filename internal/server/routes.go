package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/polo871209/go-playground/internal/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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

	apiRoute := chi.NewRouter()
	r.Mount("/api", apiRoute)

	apiRoute.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s:%s/api/swagger/doc.json", os.Getenv("HOST"), os.Getenv("PORT"))), //The url pointing to API definition
	))
	apiRoute.Get("/", s.HelloWorldHandler)
	// user routes
	apiRoute.Route("/users", func(r chi.Router) {
		r.Post("/", s.createUser)
	})

	// Auth routes
	apiRoute.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}/callback", s.getAuthCallbackHandler)
		r.Get("/{provider}", s.authHandler)
		r.Get("/logout/{provider}", s.LogoutHandler)
	})

	return r
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

// HelloWorldHandler godoc
// @Summary hello wrold
// @Tags default
// @Accept  json
// @Produce  json
// @Success 200 {object} helloWorldResponse
// @Failure 400 {object} errorResponse
// @Router /api [get]
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	err := s.writeJSON(w, http.StatusOK, helloWorldResponse{"Hello World"})
	if err != nil {
		s.errorJSON(w, err)
	}
}
