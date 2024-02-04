package inmemory

import (
	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/repository"
	"github.com/ljones140/simple_go_api/internal/repository/models"
)

// Warning this repository is not thread safe as it uses a map
func New() *Repository {
	return &Repository{
		objects: make(map[uuid.UUID]models.Object),
	}
}

type Repository struct {
	objects map[uuid.UUID]models.Object
}

func (r *Repository) InsertObject(object models.Object) (*models.Object, error) {
	object.ID = uuid.New()
	r.objects[object.ID] = object
	return &object, nil
}

func (r *Repository) GetObject(id uuid.UUID) (*models.Object, error) {
	obj, ok := r.objects[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &obj, nil
}
