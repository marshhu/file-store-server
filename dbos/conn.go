package dbos

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marshhu/file-store-server/conf"
	"log"
)
var (
	dbConn *sql.DB
	err error
	)

var conStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",conf.DBSetting.User,conf.DBSetting.Password,conf.DBSetting.Host,conf.DBSetting.DBName)

//var conStr ="root:Hdd123456@tcp(www.dooyar.com:3306)/fileserver?charset=utf8"
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
