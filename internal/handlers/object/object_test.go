package object_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ljones140/simple_go_api/internal/handlers"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestCreate_Success(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)

	payload := `{"name": "myobject"}`
	buf := bytes.NewBufferString(payload)
	req, err := http.NewRequest("POST", svr.URL+"/objects", buf)
	require.NoError(t, err)

	// Act
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func NewTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	r := chi.NewRouter()

	handlers.RegisterRoutes(r)
	svr := httptest.NewServer(r)
	t.Cleanup(svr.Close)
	return svr
}
