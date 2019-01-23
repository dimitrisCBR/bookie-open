package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
)

type MongoConfig struct {
	Address string
	Port string
	Dbname string
}

var config MongoConfig

var db * mgo.Database

func init() {

	LoadMongoConfig("./mongodatabase-conf.json")

	s, err := mgo.Dial("mongodb://"+config.Address+":"+config.Port)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db = s.DB(config.Dbname)
}

func LoadMongoConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}
}

func getDatabase() *mgo.Database {

	return db
}

func GetDB() * mgo.Database {
	return db
}