package main

import (
	"log"
	"net/http"

	"github.com/jahio/af/api"
	"github.com/jahio/af/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("af.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err.Error())
	}

	models.Init(db)
	router := api.Init(db)

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	server.ListenAndServe()

}
