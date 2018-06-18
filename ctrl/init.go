package ctrl

import (
	"github.com/jmoiron/sqlx"
)

func ConnectDB(dbName string)(db *sqlx.DB,err error){
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/"+ dbName +"?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql",dsn)
	if err != nil {
		return  nil, err
	}

	return db, err
}

var headerKeys = make(map[string]interface{})

func setHeader(){

	headerKeys = map[string]interface{}{
		"Server":"Sys_API",
		"Host":"nopadol.net:9000",
		"Content_Type":"application/json",
		"Access-Control-Allow-Origin":"*",
		"Access-Control-Allow-Methods":"GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers":"Origin, Content-Type, X-Auth-Token",
	}
}