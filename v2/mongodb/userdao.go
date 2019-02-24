package mongodb

import (
	"dimitrisCBR/bookie-open/v2/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userDao struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Username     string
	PasswordHash string
	Salt         string
}

func userDaoIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUser(u *model.User) (*userDao, error) {
	user := userDao{Username: u.Username}
	err := user.setSaltedPassword(u.Password)
	return &user, err
}

func (u *userDao) comparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *userDao) setSaltedPassword(password string) error {
	salt := uuid.New().String()
	passwordBytes := []byte(password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash[:])
	u.Salt = salt

	return nil
}
