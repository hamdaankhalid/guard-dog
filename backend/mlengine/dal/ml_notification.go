package dal

import "github.com/hamdaankhalid/mlengine/database"

func UploadMlNotification(mlNotification *MlNotification) error {
	conn, err := database.OpenConnection()
	if err != nil {
		return err
	}

	query := "NSERT INTO " + database.MlNotificationTable + " VALUES ($1, $2, $3,)"
	// TODO: insert query and variables
	_, err = conn.Exec(query, mlNotification.Id)

	return err
}

func RetrieveMlNotifications(userId int) ([]MlNotification, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return []MlNotification{}, err
	}

	var notificatons []MlNotification
	query := "SELECT * FROM " + database.MlNotificationTable + " WHERE user_id=$1"
	err = conn.Get(&notificatons, query, userId)

	return notificatons, err
}
