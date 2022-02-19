package uuid

import "github.com/google/uuid"

func Generate() string {
	return uuid.NewString()
}
