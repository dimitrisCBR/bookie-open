package mongodb

import (
	"dimitrisCBR/bookie-open/v2/config"
	"dimitrisCBR/bookie-open/v2/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AppointmentService struct {
	collection *mgo.Collection
}

var collectionAppointment = "appointment"

func NewAppointmentService(session *mgo.Session) *AppointmentService {
	collection := session.DB(config.Configuration().MongoConfig.Dbname).C(collectionAppointment)
	return &AppointmentService{collection}
}

func (as *AppointmentService) CreateAppointment(a *model.Appointment, u *model.User) error {
	appointment := newAppointmentModel(a, u)
	return as.collection.Insert(appointment)
}

func (as *AppointmentService) GetAppointmentsForUser(userId string) (error, []model.Appointment) {
	var dbResults []appointmentModel
	err := as.collection.Find(bson.M{"userid": userId}).All(&dbResults)
	modelResults := make([]model.Appointment, len(dbResults))
	for i := 0; i <= len(dbResults)-1; i++ {
		var dbAppointment = dbResults[i]
		modelResults[i] = model.Appointment{
			Id:          dbAppointment.Id.Hex(),
			Name:        dbAppointment.Name,
			Description: dbAppointment.Description,
			Fee:         dbAppointment.Fee,
			Paid:        dbAppointment.Paid,
			StartDate:   dbAppointment.StartDate,
			EndDate:     dbAppointment.EndDate}
	}
	return err, modelResults
}
