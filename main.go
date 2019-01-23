package main

import (
	"dimitrisCBR/bookie-api/v2/config"
	"dimitrisCBR/bookie-api/v2/mongodb"
	"dimitrisCBR/bookie-api/v2/server"
	"fmt"
	"log"
)

type App struct {
	server  *server.Server
	session *mongodb.Session
}

func (a *App) Initialize() {

	var err error
	a.session, err = mongodb.NewSession()
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	u := mongodb.NewUserService(a.session.Copy())
	as := mongodb.NewAppointmentService(a.session.Copy())
	a.server = server.NewServer(*u,*as, config.Configuration())
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
