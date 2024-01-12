package discussions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	discussions "github.com/CVWO/sample-go-app/internal/dataaccess"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
	"github.com/pkg/errors"
)


const (
	Listdiscussions = "discussions.Accessdiscussion"

	SuccessfulListdiscussionsMessage = "Successfully listed discussions"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrievediscussions           = "Failed to retrieve discussions in %s"
	ErrEncodeView              = "Failed to retrieve discussions in %s"
)
func AccessDiscussion(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "GetDicussion"))
	}

	fmt.Printf("to call discussion.List(db)\n")
	discussion, err := discussions.ListDiscussion(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievediscussions, "GetDicussion"))
	}

	data, err := json.Marshal(discussion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "GetDicussion"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListdiscussionsMessage},
	}, nil
}


func EditDiscussion(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
		fmt.Printf("Invalid input for edit discussion")
		return nil, errors.Wrap(err, fmt.Sprintf("Error to edit discussion %s", "editDiscussion"))
	}

	dis, err := discussions.EditDiscussion(db, discussion)
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



func CreateDiscussion(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf("Error to add discussion %s", "CreateDiscussion"))
	}

	fmt.Printf("to call CreateDiscussion(db)\n")

	var discussion models.Discussion
	err = json.NewDecoder(r.Body).Decode(&discussion)
	if err != nil {
		fmt.Printf("Invalid input for CreateDiscussion")
		return nil, errors.Wrap(err, fmt.Sprintf("Error to add discussion %s", "CreateDiscussion"))
	}

	discussions, err := discussions.CreateDiscussion(db, discussion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error to add discussion %s", "CreateDiscussion"))
	}

	data, err := json.Marshal(discussions)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrive discussion %s", "CreateDiscussion"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"successfully added a discussion"},
	}, nil
}


func AddComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "addComment"))
	}

	fmt.Printf("to call addComment(db)\n")

	var comment models.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		fmt.Printf("Invalid input for add comment\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "addComment"))
	}

	newComment, err := discussions.AddComment(db, comment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievediscussions, "addComment"))
	}

	data, err := json.Marshal(newComment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "addComment"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListdiscussionsMessage},
	}, nil
}


func LikesInc(w http.ResponseWriter, r *http.Request)(*api.Response, error){
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "LikesInc"))
	}

	var id models.DiscussionId
	err = json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		fmt.Printf("Invalid input for LikesInc\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "addComment"))
	}

	fmt.Printf("to call IncreaseLikes(db)\n")

	err = discussions.IncreaseLikes(db, id.DiscussionId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievediscussions, "LikesInc"))
	}

	
	return &api.Response{
		Payload: api.Payload{
			Data: json.RawMessage{},
		},
		Messages: []string{"Successfully liked"},
	}, nil
}


func DeleteDiscussion(w http.ResponseWriter, r *http.Request)(*api.Response, error){
	fmt.Printf("calling database.GetDB()\n")
	db, err := database.GetDB()


	if err != nil {
		fmt.Printf("error to connect to DB\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "DeleteDiscussion"))
	}

	var id models.DiscussionId
	err = json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		fmt.Printf("Invalid input for DeleteDiscussion\n")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "DeleteDiscussion"))
	}

	fmt.Printf("to call DeleteDiscussion(db)\n")

	err = discussions.DeleteDiscussion(db, id.DiscussionId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievediscussions, "DeleteDiscussion"))
	}

	
	return &api.Response{
		Payload: api.Payload{
			Data: json.RawMessage{},
		},
		Messages: []string{"Successfully deleted"},
	}, nil
}

