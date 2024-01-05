package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	users "github.com/CVWO/sample-go-app/internal/dataaccess"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
	"github.com/pkg/errors"
)

const (
	ListUsers = "users.AccessUser"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func AddUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "adduser"))
	}

	fmt.Printf("to call addUser(db)\n")

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Invalid input for add user")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "adduser"))
	}

	users, err := users.Add(db, user)
	if err != nil {
		fmt.Printf("construct error response when fail to add new user\n")
		return &api.Response{
			Payload: api.Payload{
				Data: json.RawMessage(err.Error()),

			},
			Messages: []string{"failed to create user"},
			ErrorCode: 1,
		},
		errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, "Adduser"))
	}

	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "adduser"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

func ValidUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "CheckUser"))
	}

	fmt.Printf("to call IsValidUser(db)\n")
	name := r.URL.Query().Get("name")
	user := models.User{
		Name: name,
	}
	u, err := users.IsValidUser(db, user)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, "CheckUser"))
	}
	data, err := json.Marshal(u)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "CheckUser"))
	}
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

