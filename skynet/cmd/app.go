package main

import (
	"fmt"
	"log"

	"skynet/pkg"
	"skynet/pkg/config"
	"skynet/pkg/mongo"
)

type App struct {
	session *mongo.Session
	config  *root.Config
}

func (app *App) Initialise() {
	app.config = config.GetConfig()

	var err error
	app.session, err = mongo.NewSession(app.config.Mongo)
	if err != nil {
		log.Fatalln("cannot connect to mongodb")
	}
}

func (app *App) Run() {
	fmt.Println("Running the skynet server")

	defer app.session.Close()
}
