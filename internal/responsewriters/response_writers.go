package responsewriters

import (
	"encoding/json"
	"net/http"
)

func WriteJSON[T any](w http.ResponseWriter, status int, data T) {
	respBody, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBody)
}
