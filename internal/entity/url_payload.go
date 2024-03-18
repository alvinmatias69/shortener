package entity

type UrlPayload struct {
	Id      int64  `db:"id"`
	Hash    string `db:"hash"`
	LongUrl string `db:"long_url"`
}
