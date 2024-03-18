package controller

import (
	"context"
	"strings"
	"time"

	"github.com/alvinmatias69/shortener/internal/constants"
	"github.com/alvinmatias69/shortener/internal/entity"
)

var (
	memo       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base int64 = 62
)

func createUrlPayload(url string) entity.UrlPayload {
	var (
		id   = time.Now().Unix()
		hash = createHash(id)
	)

	return entity.UrlPayload{
		Id:      id,
		Hash:    hash,
		LongUrl: url,
	}
}

func createHash(id int64) string {
	if id == 0 {
		return ""
	}

	remainder := id % base
	return createHash(id/base) + memo[remainder:remainder+1]
}

func createUrl(ctx context.Context, payload entity.UrlPayload) string {
	return strings.Join([]string{ctx.Value(constants.HostKey).(string), payload.Hash}, "/")
}
