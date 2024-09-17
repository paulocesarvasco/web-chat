package authorization

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paulocesarvasco/web-chat/authorization/internal/database"
)

type API interface {
	ValidateCredentials(w http.ResponseWriter, r *http.Request)
	Engine() *chi.Mux
}

type api struct {
	auth
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
		auth: authorizationService,
	}
	r := chi.NewRouter()
	r.Post("/login", api.ValidateCredentials)
	api.Mux = r

	return &api
}

func (a *api) Engine() *chi.Mux {
	return a.Mux
}
