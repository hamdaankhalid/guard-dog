package dal

import (
	"log"

	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/database"
	"github.com/jmoiron/sqlx"
)



type Queries struct {
	Conn *sqlx.DB
}

func (q *Queries) UploadMlNotification(mlNotification *MlNotification) error {
	query := "INSERT INTO " + database.MlNotificationTable + " VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := q.Conn.Exec(query,
		mlNotification.Id,
		mlNotification.DeviceName,
		mlNotification.SessionId,
		mlNotification.Part,
		mlNotification.ModelId,
		mlNotification.UserId,
	)

	return err
}

func (q *Queries) RetrieveAllMlNotifications(userId int) ([]MlNotification, error) {
	var notificatons []MlNotification
	query := "SELECT * FROM " + database.MlNotificationTable + " WHERE user_id=$1"
	err := q.Conn.Get(&notificatons, query, userId)

	return notificatons, err
}

func (q *Queries) RetrieveMlNotification(notificationId uuid.UUID) (MlNotification, error) {
	var notificaton MlNotification
	query := "SELECT * FROM " + database.MlNotificationTable + " WHERE id=$1"
	err := q.Conn.Get(&notificaton, query, notificationId)

	return notificaton, err
}

func (q *Queries) UploadModel(model *Model) error {
	query := "INSERT INTO " + database.ModelTable + " VALUES ($1, $2, $3, $4)"

	_, err := q.Conn.Exec(query,
		model.Id,
		model.UserId,
		model.ModelFile,
		model.Filename,
	)

	return err
}

func (q *Queries) RetrieveModel(id uuid.UUID) (Model, error) {
	query := "SELECT * FROM " + database.ModelTable + " WHERE id=$1"

	var model Model
	err := q.Conn.Get(&model, query, id)
	if err != nil {
		log.Println("Error in getting model from db")
		return Model{}, err
	}
	return model, nil
}

func (q *Queries) RetrieveAllModels(userId int) ([]ModelWithoutData, error) {
	query := "SELECT id, filename, user_id FROM " + database.ModelTable + " WHERE user_id=$1"

	var models []ModelWithoutData
	err := q.Conn.Select(&models, query, userId)
	if err != nil {
		log.Println("Error in getting models from db")
		return []ModelWithoutData{}, err
	}
	return models, nil
}

func (q *Queries) DeleteModel(modelId uuid.UUID) error {
	query := "DELETE * FROM " + database.ModelTable + " WHERE id=$1"

	var models []Model
	err := q.Conn.Select(&models, query, modelId)
	if err != nil {
		log.Println("Error in deleting model from db")
		return err
	}
	return nil
}
