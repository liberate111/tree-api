package controllers

import (
	"github.com/google/uuid"
)

func GenUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
