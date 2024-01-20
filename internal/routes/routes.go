package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/discussions"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	//create user
	return func(r chi.Router) {
		r.Post("/users", func(w http.ResponseWriter, req *http.Request) {
			response, err := users.AddUser(w, req)

			if err != nil {
				fmt.Printf("error returned for add new user: %s\n", response.Payload.Data)
				fmt.Printf("errorcode returned for add new user: %d\n", response.ErrorCode)
				fmt.Printf("error message returned for add new user: %s\n", response.Messages[0])
				b, _ := json.Marshal(response)
				var prettyJSON bytes.Buffer
				json.Indent(&prettyJSON, b, "  ", "\t")
				fmt.Println(prettyJSON)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//authenticate user
		r.Get("/checkuser", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.ValidUser(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//edit discusssion
		r.Post("/discussion", func(w http.ResponseWriter, req *http.Request) {
			response, _ := discussions.EditDiscussion(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		// get all discussions
		r.Get("/discussions", func(w http.ResponseWriter, req *http.Request) {
			response, _ := discussions.AccessDiscussion(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//add discussion/post
		r.Put("/discussion", func(w http.ResponseWriter, req *http.Request) {
			//TODO change the code below
			response, _ := discussions.CreateDiscussion(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//add comment
		r.Put("/comment", func(w http.ResponseWriter, req *http.Request) {
			//TODO change the function
			response, _ := discussions.AddComment(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//likes increase
		r.Put("/likes", func(w http.ResponseWriter, req *http.Request) {
			response, _ := discussions.LikesInc(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		//delete discussion
		r.Delete("/discussion", func(w http.ResponseWriter, req *http.Request) {
			response, _ := discussions.DeleteDiscussion(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		
	}
}

