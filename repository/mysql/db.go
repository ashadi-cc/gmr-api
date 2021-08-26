package mysql

import (
	"api-gmr/config"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbSync sync.Once

func connect() *sql.DB {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetApp().DbUser,
		config.GetApp().DbPassword,
		config.GetApp().DbHost,
		config.GetApp().DbPort,
		config.GetApp().DbName,
	)
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func getDB() *sql.DB {
	dbSync.Do(func() {
		db = connect()
	})
	return db
}
