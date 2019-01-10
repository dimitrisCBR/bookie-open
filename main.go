package main

import (
	"github.com/gorilla/mux"
	"dimitrisCBR/bookie-api/app"
	"os"
	"fmt"
	"net/http"
	"dimitrisCBR/bookie-api/controller"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controller.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controller.Authenticate).Methods("POST")
	//router.HandleFunc("/api/contacts/new", controller.CreateContact).Methods("POST")
	//router.HandleFunc("/api/me/contacts", controller.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
