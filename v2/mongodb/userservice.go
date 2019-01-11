package mongodb

import (
	"dimitrisCBR/bookie-api/v2/config"
	"dimitrisCBR/bookie-api/v2/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

var CollectionName = "user"

func NewUserService(session *mgo.Session) *UserService {
	collection := session.DB(config.Configuration().MongoConfig.Dbname).C(CollectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (p *UserService) CreateUser(u *model.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}

	return p.collection.Insert(&user)
}

func (p *UserService) GetUserByUsername(username string) (error, model.User) {
	usermodel := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&usermodel)
	return err, model.User{
		Id:       usermodel.Id.Hex(),
		Username: usermodel.Username,
		Password: "-"}
}

func (p *UserService) Login(c model.Credentials) (error, model.User) {
	usermodel := userModel{}
	err := p.collection.Find(bson.M{"username": c.Username}).One(&usermodel)

	err = usermodel.comparePassword(c.Password)
	if err != nil {
		return err, model.User{}
	}

	return err, model.User{
		Id:       usermodel.Id.Hex(),
		Username: usermodel.Username,
		Password: "-"}
}
