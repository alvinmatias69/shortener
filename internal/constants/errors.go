package constants

import "errors"

var (
	UrlNotFound = errors.New("URL not found")
	InvalidUrl  = errors.New("Invalid URL")
)
