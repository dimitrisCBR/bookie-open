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

func (as *AppointmentService) GetAppointmentsForUser(userId *string) (error, []model.Appointment) {
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

func (as *AppointmentService) FindAppointmentById(userId string, apptId string) (error, model.Appointment) {
	appModel := appointmentModel{}
	err := as.collection.Find(bson.M{"userid": userId, "_id": bson.ObjectIdHex(apptId)}).One(&appModel)
	return err, model.Appointment{
		Id:          appModel.Id.Hex(),
		Name:        appModel.Name,
		Description: appModel.Description,
		Fee:         appModel.Fee,
		Paid:        appModel.Paid,
		StartDate:   appModel.StartDate,
		EndDate:     appModel.EndDate}
}

func (as *AppointmentService) DeleteAppointmentById(userId string, appointmentId string) error {
	return as.collection.Remove(bson.M{"_id": bson.ObjectIdHex(appointmentId), "userid": userId})
}

func (as *AppointmentService) UpdateAppointment(updatedAppt *model.Appointment, userid string, ) error {
	return as.collection.Update(bson.M{"userid": userid}, bson.M{"$set":bson.M{
		"name":        updatedAppt.Name,
		"description": updatedAppt.Description,
		"fee":         updatedAppt.Fee,
		"paid":        updatedAppt.Paid,
		"startdate":   updatedAppt.StartDate,
		"enddate":     updatedAppt.EndDate}})
}
