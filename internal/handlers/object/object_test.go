package object_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ljones140/simple_go_api/internal/handlers"
	"github.com/ljones140/simple_go_api/internal/handlers/object"
	"github.com/ljones140/simple_go_api/internal/repository/inmemory"
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

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var responseData object.ObjectResponse
	require.NoError(t, json.Unmarshal(resBody, &responseData))

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
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreate_ThenGet(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)

	payload, err := json.Marshal(object.ObjectCreateRequest{Name: "test"})
	require.NoError(t, err)

	req, err := http.NewRequest("POST", svr.URL+"/objects", bytes.NewReader(payload))
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	resBody, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	var createResponseData object.ObjectResponse
	require.NoError(t, json.Unmarshal(resBody, &createResponseData))

	getReq, err := http.NewRequest("GET", svr.URL+"/objects/"+createResponseData.ID, nil)
	require.NoError(t, err)

	getRes, err := http.DefaultClient.Do(getReq)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRes.StatusCode)

	getResBody, err := io.ReadAll(getRes.Body)
	require.NoError(t, err)
	var responseData object.ObjectResponse
	require.NoError(t, json.Unmarshal(getResBody, &responseData))

	assert.Equal(t, "test", responseData.Name)
	assert.Equal(t, createResponseData.ID, responseData.ID)
}

func TestGet_ObjectNotFound(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)
	req, err := http.NewRequest("GET", svr.URL+fmt.Sprintf("/objects/%s", uuid.NewString()), nil)
	require.NoError(t, err)
	// Act
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestGet_Object_InvalidUUID(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)
	req, err := http.NewRequest("GET", svr.URL+"/objects/invalid-uuid", nil)
	require.NoError(t, err)
	// Act
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	// Assert
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func NewTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	r := chi.NewRouter()
	handlers.RegisterRoutes(r, inmemory.New())

	svr := httptest.NewServer(r)
	t.Cleanup(svr.Close)
	return svr
}
