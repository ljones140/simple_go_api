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

	req, err := http.NewRequest(
		"POST", svr.URL+"/objects", ModelToReader(t, object.ObjectCreateRequest{Name: "test"}),
	)
	require.NoError(t, err)

	// Act
	var responseData object.ObjectResponse
	MakeRequest(t, req, http.StatusCreated, &responseData)

	// Assert
	assert.Equal(t, "test", responseData.Name)
}

func TestCreate_ErrorNonJsonBody(t *testing.T) {
	// Arrange
	svr := NewTestServer(t)
	req, err := http.NewRequest("POST", svr.URL+"/objects", bytes.NewReader([]byte("not json")))
	require.NoError(t, err)

	// Act
	MakeRequest(t, req, http.StatusBadRequest, nil)
}

func TestCreate_ThenGet(t *testing.T) {
	svr := NewTestServer(t)

	req, err := http.NewRequest(
		"POST", svr.URL+"/objects", ModelToReader(t, object.ObjectCreateRequest{Name: "test"}),
	)
	require.NoError(t, err)

	var createResponse object.ObjectResponse
	MakeRequest(t, req, http.StatusCreated, &createResponse)

	getReq, err := http.NewRequest("GET", svr.URL+"/objects/"+createResponse.ID, nil)
	require.NoError(t, err)

	var getResponse object.ObjectResponse
	MakeRequest(t, getReq, http.StatusOK, &getResponse)

	assert.Equal(t, "test", getResponse.Name)
	assert.Equal(t, createResponse.ID, getResponse.ID)
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
	MakeRequest(t, req, http.StatusNotFound, nil)
}

func NewTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	r := chi.NewRouter()
	handlers.RegisterRoutes(r, inmemory.New())

	svr := httptest.NewServer(r)
	t.Cleanup(svr.Close)
	return svr
}

func MakeRequest(t *testing.T, req *http.Request, wantStatus int, resObject any) {
	t.Helper()
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	assert.Equal(t, wantStatus, res.StatusCode)

	if resObject != nil {
		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(resBody, resObject))
	}
}

func ModelToReader(t *testing.T, model interface{}) io.Reader {
	t.Helper()
	payload, err := json.Marshal(model)
	require.NoError(t, err)
	return bytes.NewReader(payload)
}
