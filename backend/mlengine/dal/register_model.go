package dal

import (
	"mime/multipart"

	"github.com/hamdaankhalid/mlengine/database"
)

type Model struct {
	Id        int
	UserId    int
	ModelFile multipart.File
}

func UploadModel(model *Model) error {
	conn, err := database.OpenConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(`INSERT INTO users VALUES ($1, $2, $3,)`,
		model.Id,
		model.UserId,
		model.ModelFile,
	)

	return err
}
