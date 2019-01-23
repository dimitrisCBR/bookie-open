package model

import (
	"dimitrisCBR/bookie-api/src/v1c/v1/db"
	u "dimitrisCBR/bookie-api/v2/v1/utiltil"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"strings"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	ID       bson.ObjectId `bson:"_id"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    string        `json:"token";sql:"-"`
}

type Accounts []Account

var COLLECTION = "Accounts"

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := db.GetDB().C(COLLECTION).Table("Accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (m *MovieModel) Create(data forms.CreateMovieCommand) error {
	collection := dbConnect.Use("test-mgo", "movies")
	err := collection.Insert(bson.M{"name": data.Name, "rating": data.Rating, "desc": data.Desc})
	return err
}

func (m *MovieModel) Find() (list []Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *MovieModel) Get(id string) (movie Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MovieModel) Update(id string, data forms.UpdateMovieCommand) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return err
}

func (m *MovieModel) Delete(id string) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := db.GetDB().C(COLLECTION).Find(bson.M{"email": email}).
		Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("Accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
