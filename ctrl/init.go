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

