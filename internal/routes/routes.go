package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.AccessUsers(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		r.Post("/users", func(w http.ResponseWriter, req *http.Request) {
			user, _ := users.AddUser(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		})
		r.Get("/checkuser", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.ValidUser(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

	}
}

