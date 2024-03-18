package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/alvinmatias69/shortener/internal/constants"
	"github.com/alvinmatias69/shortener/internal/entity"
)

type Handler struct {
	controller controller
}

func New(controller controller) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) GetShortened(res http.ResponseWriter, req *http.Request) {
	var (
		url            = req.PathValue("url")
		ctx, cancelCtx = context.WithTimeout(context.Background(), time.Second*60)
	)
	defer cancelCtx()

	longUrl, err := h.controller.Get(ctx, url)
	if errors.Is(err, constants.UrlNotFound) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("URL not found"))
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("error please try again"))
		return
	}

	http.Redirect(res, req, longUrl, http.StatusMovedPermanently)
}

func (h *Handler) CreateShortened(res http.ResponseWriter, req *http.Request) {
	var (
		requestBody    entity.CreateRequest
		ctx, cancelCtx = context.WithTimeout(context.Background(), time.Second*60)
	)
	defer cancelCtx()

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request body"))
	}

	ctxWithHost := context.WithValue(ctx, constants.HostKey, req.Host)
	responseBody, err := h.controller.Create(ctxWithHost, requestBody)
	if errors.Is(err, constants.InvalidUrl) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid URL"))
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("error please try again"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(responseBody)
}
