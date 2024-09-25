package authorization

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paulocesarvasco/web-chat/authorization/internal/auth0"
	"github.com/paulocesarvasco/web-chat/authorization/internal/database"
)

type API interface {
	CreateClient() http.HandlerFunc
	Engine() *chi.Mux
	ValidateCredentials() http.HandlerFunc
	ValidateToken() http.HandlerFunc
}

type api struct {
	auth
	auth0.API
	*chi.Mux
}

func NewAPI() API {
	databaseClient, err := database.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	authorizationService := auth{
		db: databaseClient,
	}
	api := api{
		API:  auth0.New(),
		auth: authorizationService,
	}
	r := chi.NewRouter()
	r.Post("/login", api.ValidateCredentials())
	r.Post("/create", api.CreateClient())
	r.Get("/validate", api.ValidateToken())
	api.Mux = r

	return &api
}

func (a *api) Engine() *chi.Mux {
	return a.Mux
}
