package mongodb

import (
	"dimitrisCBR/bookie-api/v2/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type appointmentModel struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	UserId      string        `json:"user_id"`
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}

func newAppointmentModel(a *model.Appointment, u *model.User) *appointmentModel {
	appointment := appointmentModel{
		UserId:      u.Id,
		Name:        a.Name,
		Description: a.Description,
		StartDate:   a.StartDate,
		EndDate:     a.EndDate}
	return &appointment
}
