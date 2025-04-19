package vo

import "github.com/oklog/ulid/v2"

type NewsID ulid.ULID

func NewNewsID(id string) NewsID {
	return NewsID(ulid.Make())
}
