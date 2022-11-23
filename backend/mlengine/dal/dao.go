package dal

import (
	"github.com/google/uuid"
)

type Model struct {
	ModelFile []byte    `db:"model_file" json:"modelFile"`
	Filename  string    `db:"filename" json:"filename"`
	Id        uuid.UUID `db:"id" json:"id"`
	UserId    int       `db:"user_id" json:"userId"`
}

type MlNotification struct {
	DeviceName string    `db:"device_name" json:"deviceName"`
	Id         uuid.UUID `db:"id" json:"id"`
	SessionId  int       `db:"session_id" json:"sessionId"`
	Part       int       `db:"part" json:"part"`
	ModelId    uuid.UUID `db:"model_id" json:"modelId"`
	UserId     int       `db:"user_id" json:"userId"`
}
