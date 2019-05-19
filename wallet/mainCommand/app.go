package main

import (
	"log"

	root "wallet/pkg"
	"wallet/pkg/config"
	"wallet/pkg/mongo"
	"wallet/pkg/server"
)

// App forms the core struct for running the site
type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

// Initialize bootstraps the app
func (a *App) Initialize() {
	a.config = config.GetConfig()

	var err error
	// check whether the wallet has been initialized before or not
	err = config.CheckBootConfigFile(a.config.Boot)
	if err != nil {
		log.Fatal("unable to create configuration file")
	}

	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatal("unable to connect to mongodb")
	}

	a.server = server.NewServer(a.config)
	a.server.CreateRoutes()
	a.server.CreateBootRouter(mongo.NewUserService(a.config.Boot.BootConfigPath), mongo.NewDeviceService(a.config.Boot.DeviceConfigPath))
	a.server.CreatePairIdentityRouter(mongo.NewPairIdentityService(a.session, a.config.Mongo))
}

// Run starts the server
func (a *App) Run() {

	defer a.session.Close()
	a.server.Start(a.config.Boot.BootConfigExists)
}

func main() {
	a := App{}
	a.Initialize()
	a.Run()
}
