package config

import "database/sql"

func MysqlConnectBuiltin() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/alta_online_shop")
	if err != nil {
		panic(err.Error())
	}
		
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db
}