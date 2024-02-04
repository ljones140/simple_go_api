package inmemory_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/repository"
	"github.com/ljones140/simple_go_api/internal/repository/inmemory"
	"github.com/ljones140/simple_go_api/internal/repository/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertObject(t *testing.T) {
	repo := inmemory.New()

	obj := models.Object{Name: "test"}

	createdObj, err := repo.InsertObject(obj)
	require.NoError(t, err)

	assert.Equal(t, obj.Name, createdObj.Name)
	assert.True(t, createdObj.ID != uuid.Nil)
}

func TestGetObject_Object_Exists(t *testing.T) {
	repo := inmemory.New()
	obj := models.Object{Name: "test"}
	createdObj, err := repo.InsertObject(obj)
	require.NoError(t, err)

	foundObj, err := repo.GetObject(createdObj.ID)
	require.NoError(t, err)
	assert.Equal(t, createdObj, foundObj)
}

func TestGetObject_Object_NotFound(t *testing.T) {
	repo := inmemory.New()
	_, err := repo.GetObject(uuid.New())
	assert.ErrorIs(t, err, repository.ErrNotFound)
}
