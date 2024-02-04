package object

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/responsewriters"
)

type Object struct{}

func New() *Object {
	return &Object{}
}

type ObjectCreateRequest struct {
	Name string `json:"name"`
}

type ObjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (o *Object) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createData := ObjectCreateRequest{}
	err = json.Unmarshal(reqBody, &createData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseData := ObjectResponse{
		ID:   uuid.NewString(),
		Name: createData.Name,
	}

	responsewriters.WriteJSON(w, http.StatusCreated, responseData)
}
