package main

import (
	"fmt"
	"log"

	"skynet/pkg"
	"skynet/pkg/config"
	"skynet/pkg/mongo"
	"skynet/pkg/server"
)

type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

func (a *App) Initialize() {
	a.config = config.GetConfig()

	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatal("unable to connect to mongodb")
	}

	u := mongo.NewUserService(a.session, a.config.Mongo)
	a.server = server.NewServer(u, a.config)
}

func (a *App) Run() {
	fmt.Println("Run")
	defer a.session.Close()
	a.server.Start()
}

func main() {
	a := App{}
	a.Initialize()
	a.Run()
}
