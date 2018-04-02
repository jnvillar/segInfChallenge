package app

import (
	"api/app/items"
	"api/app/index"
	"api/app/files"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	// Needed to sql lite 3
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"io/ioutil"
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
	//db, err := sql.Open("sqlite3", "./foo.db")


	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", "user", "userpwd", "db:3306", "db"))
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
		itemsQuery, err := ioutil.ReadFile("./src/api/items.sql")
		if err != nil {panic("Could not read items.sql")}

		filesQuery , err := ioutil.ReadFile("./src/api/files.sql")
		if err != nil {panic("Could not read files.sql")}

		createTable(db, string(itemsQuery))
		createTable(db, string(filesQuery))
		
		return db
	}

}

func createTable(db *sql.DB, table string) {
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}
