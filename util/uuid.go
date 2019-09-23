package util

import (
	uuid "github.com/satori/go.uuid"
)


func NewUUID() string{
	id,_ := uuid.NewV4()
    return id.String()
}
