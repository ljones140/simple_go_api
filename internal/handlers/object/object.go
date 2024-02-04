package object

import "net/http"

type Object struct{}

func New() *Object {
	return &Object{}
}

func (o *Object) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
