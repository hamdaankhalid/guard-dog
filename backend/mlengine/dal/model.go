package dal

import "github.com/hamdaankhalid/mlengine/database"

func UploadModel(model *Model) error {
	conn, err := database.OpenConnection()
	if err != nil {
		return err
	}

	query := "INSERT INTO " + database.ModelTable + " VALUES ($1, $2, $3, $4)"
	bytedata := []byte{}
	_, err = model.ModelFile.Read(bytedata)

	if err != nil {
		return err
	}

	_, err = conn.Exec(query,
		model.Id,
		model.UserId,
		bytedata,
		model.Filename,
	)

	return err
}

func RetrieveModel(id int) (Model, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return Model{}, err
	}

	query := "SELECT * FROM " + database.ModelTable + " WHERE id=$1"

	var model Model
	err = conn.Get(&model, query, id)

	return model, err
}

func RetrieveAllModels(userId int) ([]Model, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return []Model{}, err
	}

	query := "SELECT * FROM " + database.ModelTable + " WHERE user_id=$1"

	var models []Model
	err = conn.Get(&models, query, userId)

	return models, err
}
