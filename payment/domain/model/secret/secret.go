package secret

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Secret struct {
	Id        int64
	AccountId int64
	Key       string
	CreatedAt time.Time
}

func NewSecret() (*Secret, error) {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	encoded := base64.URLEncoding.EncodeToString(key)
	return &Secret{Key: encoded}, nil
}
