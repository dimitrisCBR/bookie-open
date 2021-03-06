package server

import (
	"dimitrisCBR/bookie-open/v2/model"
	"dimitrisCBR/bookie-open/v2/mongodb"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type appointmentRouter struct {
	appointmentService mongodb.AppointmentService
	userService        mongodb.UserService
	authHelper         *authHelper
}

func NewAppointmentRouter(as mongodb.AppointmentService, us mongodb.UserService, router *mux.Router, a *authHelper) *mux.Router {
	appointmentRouter := appointmentRouter{as, us, a}
	router.HandleFunc("/get", a.validate(appointmentRouter.getAppointmentsHandler)).Methods("GET")
	router.HandleFunc("/create", a.validate(appointmentRouter.createAppointmentHandler)).Methods("POST")
	router.HandleFunc("/update", a.validate(appointmentRouter.updateAppointmentHandler)).Methods("PUT")
	router.HandleFunc("/delete", a.validate(appointmentRouter.deleteAppointmentHandler)).Methods("DELETE")
	return router
}

func (ar *appointmentRouter) createAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
	if !ok {
		Error(w, http.StatusBadRequest, "no context")
		return
	}
	username := claim.Username

	err, user := ar.userService.GetUserByUsername(&username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	err, appointment := decodeAppointment(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ar.appointmentService.CreateAppointment(&appointment, &user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, appointment)
}

func (ar *appointmentRouter) updateAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
	if !ok {
		Error(w, http.StatusBadRequest, "no context")
		return
	}
	username := claim.Username

	err, user := ar.userService.GetUserByUsername(&username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	err, appointment := decodeAppointment(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ar.appointmentService.UpdateAppointment(&appointment, user.Id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, appointment)
}

func (ar *appointmentRouter) deleteAppointmentHandler(w http.ResponseWriter, r *http.Request) {

	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
	if !ok {
		Error(w, http.StatusBadRequest, "no context")
		return
	}
	username := claim.Username

	err, user := ar.userService.GetUserByUsername(&username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	appId := r.URL.Query().Get("id")
	err, _ = ar.appointmentService.FindAppointmentById(user.Id, appId)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid appointment Id")
		return
	}

	err = ar.appointmentService.DeleteAppointmentById(user.Id, appId)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, nil)
}

func (ar *appointmentRouter) getAppointmentsHandler(w http.ResponseWriter, r *http.Request) {

	claim, ok := r.Context().Value(contextKeyAuthtoken).(claims)
	if !ok {
		Error(w, http.StatusBadRequest, "no context")
		return
	}

	username := claim.Username

	err, user := ar.userService.GetUserByUsername(&username)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	err, appointments := ar.appointmentService.GetAppointmentsForUser(&user.Id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, appointments)
}

func decodeAppointment(r *http.Request) (error, model.Appointment) {
	var a model.Appointment
	if r.Body == nil {
		return errors.New("no request body"), a
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	return err, a
}
