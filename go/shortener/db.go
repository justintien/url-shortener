package shortener

import (
	"database/sql"
	"fmt"
	"log"

	// Add mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//InitDB inits mysql database
func InitDB(datasource, dbUser, dbPass, dbName, port string) *sql.DB {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s", dbUser, dbPass, dbName, port)

	db, err := sql.Open(datasource, connectionString)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
