package mongodb

import (
	"dimitrisCBR/bookie-open/v2/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type appointmentModel struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	UserId      string
	Name        string
	Description string
	Fee         float32
	Paid        bool
	StartDate   time.Time
	EndDate     time.Time
}

func newAppointmentModel(a *model.Appointment, u *model.User) *appointmentModel {
	appointment := appointmentModel{
		UserId:      u.Id,
		Name:        a.Name,
		Description: a.Description,
		Fee:         a.Fee,
		Paid:        a.Paid,
		StartDate:   a.StartDate,
		EndDate:     a.EndDate}
	return &appointment
}
