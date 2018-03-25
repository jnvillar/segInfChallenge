package app

import (
	"api/app/items"
	"api/app/index"
	"api/app/files"
	"database/sql"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	// Needed to sql lite 3
	_ "github.com/mattn/go-sqlite3"
)

var (
	r *gin.Engine
)

const (
	port string = ":8080"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	index.Configure(r)
	items.Configure(r, db)
	files.Configure(r, db)
	r.Run(port)
}

func configDataBase() *sql.DB {
	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "user", "userpwd", "db", "db"))
	if err != nil {
		panic("Could not connect to the db")
	}

	for {
		err := db.Ping()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		// This is bad practice... You should create a schema.sql with all the definitions
		itemsQuery := `
				CREATE TABLE IF NOT EXISTS items(
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					name TEXT,
					description TEXT
				);`

		filesQuery := `
				CREATE TABLE IF NOT EXISTS files(
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					title TEXT,
					description TEXT
				);`

		createTable(db, itemsQuery)
		createTable(db, filesQuery)
		return db
	}

}

func createTable(db *sql.DB, table string) {
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}
