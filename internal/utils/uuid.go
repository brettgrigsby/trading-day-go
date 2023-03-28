package utils

import (
	"github.com/google/uuid"
)

func CreateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}