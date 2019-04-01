package middleware


import (
	"github.com/satori/go.uuid"
)

//生成uuid
func CreateUuid() uuid.UUID {
	u := uuid.Must(uuid.NewV4())

	return u
}



