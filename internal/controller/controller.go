package controller

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/alvinmatias69/shortener/internal/constants"
	"github.com/alvinmatias69/shortener/internal/entity"
)

type Controller struct {
	repository repository
}

func New(repository repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) Get(ctx context.Context, hash string) (string, error) {
	urlPayload, err := c.repository.Get(ctx, hash)
	if err != nil {
		return "", err
	}

	return urlPayload.LongUrl, nil
}

func (c *Controller) Create(ctx context.Context, req entity.CreateRequest) (entity.CreateResponse, error) {
	_, err := url.ParseRequestURI(req.Url)
	if err != nil {
		return entity.CreateResponse{}, constants.InvalidUrl
	}

	current, err := c.repository.GetByLongUrl(ctx, req.Url)
	if err == nil {
		return entity.CreateResponse{
			Url: createUrl(ctx, current),
		}, nil
	}

	if err != nil && !errors.Is(err, constants.UrlNotFound) {
		fmt.Printf("error while fetching: %v\n", err)
	}

	payload := createUrlPayload(req.Url)
	err = c.repository.Create(ctx, payload)
	if err != nil {
		return entity.CreateResponse{}, err
	}

	return entity.CreateResponse{
		Url: createUrl(ctx, payload),
	}, nil
}
