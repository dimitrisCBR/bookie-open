package mongodb

import (
	"dimitrisCBR/bookie-open/v2/config"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

type Session struct {
	session *mgo.Session
}

func NewSession() (*Session, error) {

	var mongoConfig = config.Configuration().MongoConfig

	info := &mgo.DialInfo{
		Addrs:    []string{mongoConfig.Address + ":" + mongoConfig.Port},
		Timeout:  30 * time.Second,
		Database: mongoConfig.Dbname,
		Username: mongoConfig.Dbuser,
		Password: mongoConfig.DbPass,
	}

	s, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &Session{s}, err
}

func (s *Session) Copy() *mgo.Session {
	return s.session.Copy()
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
