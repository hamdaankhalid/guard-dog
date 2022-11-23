package database

import "github.com/jmoiron/sqlx"

var (
	ModelTable          = "model"
	MlNotificationTable = "ml_notification"
)

func InitDb() error {
	conn, err := OpenConnection()
	if err != nil {
		return err
	}
	createTables(conn)
	return nil
}

func createTables(conn *sqlx.DB) {
	var schema = `
		CREATE TABLE model (
				id integer
		);

		CREATE TABLE ml_notification (
				id integer
		)`

	conn.MustExec(schema)
}
