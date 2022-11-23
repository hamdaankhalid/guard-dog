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
				id UUID PRIMARY KEY,
				user_id numeric,
				model_file bytea
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
