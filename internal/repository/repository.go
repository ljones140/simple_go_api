package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/repository/models"
)

var ErrNotFound = errors.New("not found")

type Repository interface {
	InsertObject(object models.Object) (*models.Object, error)
	GetObject(id uuid.UUID) (*models.Object, error)
}
