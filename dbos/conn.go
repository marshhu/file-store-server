package dbos

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
const(
	Host string =   "www.dooyar.com"
	Port int =   3306
	User string =   "root"
	Pwd string =   "Hdd123456"
	DbName string = "fileserver"
)
var (
	dbConn *sql.DB
	err error
	)

var conStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",User,Pwd,Host, Port,DbName)

func init(){
	dbConn, err = sql.Open("mysql", conStr)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// Open doesn't open a connection. Validate DSN data:
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
