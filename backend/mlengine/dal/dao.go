package dal

import "mime/multipart"

type Model struct {
	ModelFile  multipart.File
	Id, UserId int
}

type MlNotification struct {
	DeviceName                   string
	Id, SessionId, Part, ModelId int
}
