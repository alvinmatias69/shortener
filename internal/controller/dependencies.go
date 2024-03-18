package controller

import (
	"context"

	"github.com/alvinmatias69/shortener/internal/entity"
)

type repository interface {
	Get(context.Context, string) (entity.UrlPayload, error)
	GetByLongUrl(context.Context, string) (entity.UrlPayload, error)
	Create(context.Context, entity.UrlPayload) error
}
