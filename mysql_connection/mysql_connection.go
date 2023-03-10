package mysql_connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type Connection struct {
	db *sql.DB
}

func NewMySQLConnection() (*Connection, error) {
	var dataSourceString = fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dataSourceString)
	if err != nil {
		return nil, err
	}

	/*
		It's required to ensure connections are closed by the driver safely before connection is closed by MySQL server, OS, or other middlewares.
		Since some middlewares close idle connections by 5 minutes, we recommend timeout shorter than 5 minutes.
		This setting helps load balancing and changing system variables too.
	*/
	db.SetConnMaxLifetime(time.Minute * 3)
	/*
		It's highly recommended to limit the number of connection used by the application.
		There is no recommended limit number because it depends on application and MySQL server.
	*/
	db.SetMaxOpenConns(10)
	/*
		It's recommended to be set same to db.SetMaxOpenConns().
		When it is smaller than SetMaxOpenConns(), connections can be opened and closed much more frequently than you expect.
		Idle connections can be closed by the db.SetConnMaxLifetime().
		If you want to close idle connections more rapidly, you can use db.SetConnMaxIdleTime() since Go 1.15.
	*/
	db.SetMaxIdleConns(10)

	return &Connection{db}, nil
}
