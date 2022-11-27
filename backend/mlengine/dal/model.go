package dal

import (
	"log"

	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/database"
)

func UploadModel(model *Model) error {
	conn, err := database.OpenConnection()
	if err != nil {
		return err
	}

	query := "INSERT INTO " + database.ModelTable + " VALUES ($1, $2, $3, $4)"

	if err != nil {
		return err
	}

	_, err = conn.Exec(query,
		model.Id,
		model.UserId,
		model.ModelFile,
		model.Filename,
	)
	if err != nil {
		log.Println("Error in inserting model in db")
		return err
	}
	return nil
}

func RetrieveModel(id uuid.UUID) (Model, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return Model{}, err
	}

	query := "SELECT * FROM " + database.ModelTable + " WHERE id=$1"

	var model Model
	err = conn.Get(&model, query, id)
	if err != nil {
		log.Println("Error in getting model from db")
		return Model{}, err
	}
	return model, nil
}

func RetrieveAllModels(userId int) ([]ModelWithoutData, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return []ModelWithoutData{}, err
	}

	query := "SELECT id, filename, user_id FROM " + database.ModelTable + " WHERE user_id=$1"

	var models []ModelWithoutData
	err = conn.Select(&models, query, userId)
	if err != nil {
		log.Println("Error in getting models from db")
		return []ModelWithoutData{}, err
	}
	return models, nil
}

func DeleteModel(modelId uuid.UUID) error {
	conn, err := database.OpenConnection()

	if err != nil {
		return err
	}

	query := "DELETE * FROM " + database.ModelTable + " WHERE id=$1"

	var models []Model
	err = conn.Select(&models, query, modelId)
	if err != nil {
		log.Println("Error in deleting model from db")
		return err
	}
	return nil
}
