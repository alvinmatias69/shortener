package server

import "net/http"

type handler interface {
	GetShortened(http.ResponseWriter, *http.Request)
	CreateShortened(http.ResponseWriter, *http.Request)
}
