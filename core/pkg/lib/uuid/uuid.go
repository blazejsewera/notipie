package uuid

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
)

func Generate() string {
	return uuid.NewString()
}

func GenerateBasedOnContent(c []byte) string {
	hash := sha256.Sum256(c)
	hashSlice := hash[:]

	return base64.StdEncoding.EncodeToString(hashSlice)
}
