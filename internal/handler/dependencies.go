package handler

import (
	"context"

	"github.com/alvinmatias69/shortener/internal/entity"
)

type controller interface {
	Get(context.Context, string) (string, error)
	Create(context.Context, entity.CreateRequest) (entity.CreateResponse, error)
}
