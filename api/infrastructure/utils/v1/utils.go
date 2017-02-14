package utils

import (
	"github.com/satori/go.uuid"
)

//NewGuid generates new Guid
func NewGuid() uuid.UUID {
	return uuid.NewV4()
}
