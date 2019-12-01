package main

import (
	"os"
	"shortener"

	"github.com/gorilla/mux"
)

func serve(a shortener.App) {
	a.Router = mux.NewRouter()
	a.Init()
	a.Run(":" + os.Getenv("GO_EXPOSED_PORT"))
}

func main() {
	a := shortener.App{}
	a.DB = shortener.InitDB(
		"mysql",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	serve(a)
}
