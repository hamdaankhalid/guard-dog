package dal

import (
	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/database"
)

func UploadMlNotification(mlNotification *MlNotification) error {
	conn, err := database.OpenConnection()
	if err != nil {
		return err
	}

	query := "INSERT INTO " + database.MlNotificationTable + " VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = conn.Exec(query,
		mlNotification.Id,
		mlNotification.DeviceName,
		mlNotification.SessionId,
		mlNotification.Part,
		mlNotification.ModelId,
		mlNotification.UserId,
	)

	return err
}

func RetrieveAllMlNotifications(userId int) ([]MlNotification, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return []MlNotification{}, err
	}

	var notificatons []MlNotification
	query := "SELECT * FROM " + database.MlNotificationTable + " WHERE user_id=$1"
	err = conn.Get(&notificatons, query, userId)

	return notificatons, err
}

func RetrieveMlNotification(notificationId uuid.UUID) (MlNotification, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return MlNotification{}, err
	}

	var notificaton MlNotification
	query := "SELECT * FROM " + database.MlNotificationTable + " WHERE id=$1"
	err = conn.Get(&notificaton, query, notificationId)

	return notificaton, err
}
