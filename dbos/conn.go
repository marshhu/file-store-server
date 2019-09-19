package dbos

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
)

func init(){
	dbConn, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = dbConn.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}