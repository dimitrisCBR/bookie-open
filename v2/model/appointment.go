package model

import "time"

type Appointment struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Fee         float32   `json:"fee"`
	Paid        bool      `json:"paid"`
	StartDate   time.Time `json:"start"`
	EndDate     time.Time `json:"end"`
}

type AppointmentService interface {
	CreateAppointment(a *Appointment) error
	GetAppointmentBy(username string) (error, User)
	Login(c Credentials) (error, User)
}
