package discussions

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
func AccessDiscussion(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "GetDicussion"))
	}

	fmt.Printf("to call discussion.List(db)\n")
	discussion, err := users.ListDiscussion(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, "GetDicussion"))
	}

	data, err := json.Marshal(discussion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "GetDicussion"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}


func Edit(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf("Error to edit discussion %s", "editDiscussion"))
	}

	fmt.Printf("to call edit(db)\n")

	var discussion models.Discussion
	err = json.NewDecoder(r.Body).Decode(&discussion)
	if err != nil {
		fmt.Printf("Invalid input for add user")
		return nil, errors.Wrap(err, fmt.Sprintf("Error to edit discussion %s", "editDiscussion"))
	}

	dis, err := users.EditDiscussion(db, discussion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error to edit discussion %s", "editDiscussion"))
	}

	data, err := json.Marshal(dis)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrive discussion %s", "editDiscussion"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"Successfully edit the discussion"},
	}, nil
}
