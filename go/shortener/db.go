package shortener

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB(datasource, dbUser, dbPass, dbName, port string) *gorm.DB {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s?parseTime=true", dbUser, dbPass, dbName, port)

	db, err := gorm.Open(datasource, connectionString)

	if err != nil {
		log.Panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
