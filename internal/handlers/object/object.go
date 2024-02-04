package object

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/repository"
	"github.com/ljones140/simple_go_api/internal/repository/models"
	"github.com/ljones140/simple_go_api/internal/responsewriters"
)

type Object struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Object {
	return &Object{repo: repo}
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

	createdObj, err := o.repo.InsertObject(models.Object{Name: createData.Name})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := ObjectResponse{
		ID:   createdObj.ID.String(),
		Name: createdObj.Name,
	}

	responsewriters.WriteJSON(w, http.StatusCreated, responseData)
}

func (o *Object) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	obj, err := o.repo.GetObject(uuid)
	if err != nil {
		panic(err)
	}
	responseData := ObjectResponse{
		ID:   obj.ID.String(),
		Name: obj.Name,
	}

	responsewriters.WriteJSON(w, http.StatusOK, responseData)
}
