package services

import (
	b64 "encoding/base64"

	"github.com/google/uuid"
)

func GetUUID() string {
	uu := uuid.Must(uuid.NewRandom())
	data := uu.NodeID()
	str := b64.StdEncoding.EncodeToString(data)
	return str[:6]
}
