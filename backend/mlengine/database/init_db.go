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
	createModelTable(conn)
	createMlNotificationTable(conn)
	return nil
}

func createModelTable(conn *sqlx.DB) {
	query := `CREATE TABLE ` + ModelTable + ` ( 
			column1 datatype,
			column2 datatype,
			column3 datatype,
			columnN datatype,
			PRIMARY KEY( one or more columns )
	 );`

	conn.MustExec(query)
}

func createMlNotificationTable(conn *sqlx.DB) {
	query := `CREATE TABLE ` + MlNotificationTable + ` (
		column1 datatype,
		column2 datatype,
		column3 datatype,
		columnN datatype,
		PRIMARY KEY( one or more columns )
 	);`

	conn.MustExec(query)
}
