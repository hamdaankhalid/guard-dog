package database

import "github.com/jmoiron/sqlx"

const (
	ModelTable          = "model"
	MlNotificationTable = "ml_notification"
)

func InitDb(recreate bool) error {
	conn, err := OpenConnection()
	if err != nil {
		return err
	}
	if recreate {
		dropTables(conn)
	}

	createTables(conn)
	return nil
}

func createTables(conn *sqlx.DB) {
	var schema = `
		CREATE TABLE model (
				id UUID PRIMARY KEY,
				user_id numeric,
				model_file bytea,
				filename text
		);

		CREATE TABLE ml_notification (
			id UUID PRIMARY KEY,
			user_id numeric,
			model_id UUID,
			device_name text,
			session_id numeric,
			part numeric,
			CONSTRAINT fk_model
          FOREIGN KEY(model_id)
    	  REFERENCES model(id)
    	  ON DELETE CASCADE
		)`

	conn.MustExec(schema)
}

func dropTables(conn *sqlx.DB) {
	schema := `
		DROP TABLE model;
		DROP TABLE ml_notification;
	`

	conn.MustExec(schema)
}
