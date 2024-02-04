package object_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ljones140/simple_go_api/internal/handlers"
	"github.com/ljones140/simple_go_api/internal/handlers/object"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestCreate_Success(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)

	requestData := object.ObjectCreateRequest{Name: "test"}

	payload, err := json.Marshal(requestData)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", svr.URL+"/objects", bytes.NewReader(payload))
	require.NoError(t, err)

	// Act
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	resBody, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	responseData := object.ObjectResponse{}

	err = json.Unmarshal(resBody, &responseData)

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, "test", responseData.Name)
}

func TestCreate_ErrorNonJsonBody(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)
	req, err := http.NewRequest("POST", svr.URL+"/objects", bytes.NewReader([]byte("not json")))
	require.NoError(t, err)
	// Act
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	// Assert
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func NewTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	r := chi.NewRouter()

	handlers.RegisterRoutes(r)
	svr := httptest.NewServer(r)
	t.Cleanup(svr.Close)
	return svr
}
