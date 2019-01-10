package model

import (
	u "dimitrisCBR/bookie-api/src/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	Id        uint      `json:"id"` //The user that this contact belongs to
	UserId    uint      `json:"user_id"`
	ContactId uint      `json: client_id`
	Date      time.Time `json:"name"`
}

func (appointment *Appointment) Validate() (map[string]interface{}, bool) {

	if (appointment.UserId == 0) {
		return u.Message(false, "Appointment has no "), false
	}

	return u.Message(true, "success"), true
}

func (appointment *Appointment) Create() (map[string]interface{}) {

	if resp, ok := appointment.Validate(); !ok {
		return resp
	}

	GetDB().Create(appointment)

	resp := u.Message(true, "success")
	resp["appointment"] = appointment
	return resp
}

func GetAppointments(userid uint) ([]*Appointment) {

	appmnts := make([]*Appointment, 0)
	err := GetDB().Table("appointments").Where("user_id = ?", userid).Find(&appmnts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return appmnts
}
