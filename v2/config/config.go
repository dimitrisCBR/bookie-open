package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MongoConfig struct {
	Address string `json:"Address"`
	Port    string `json:"Port"`
	Dbname  string `json:"Dbname"`
	Dbuser  string `json: Dbuser`
	DbPass  string `json: Dbpass`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type AuthConfig struct {
	Secret string `json:"secret"`
}

type Config struct {
	MongoConfig  MongoConfig
	ServerConfig ServerConfig
	AuthConfig   AuthConfig
}

var signKey []byte

var config Config

func init() {
	var moConfig = LoadMongoConfig("./mongo_conf.json")
	var siConfig = LoadSecurityConfig("./security_conf.json")
	var seConfig = LoadServerConfig("./server_conf.json")
	config = Config{
		MongoConfig:  moConfig,
		ServerConfig: seConfig,
		AuthConfig:   siConfig}
}

func LoadMongoConfig(path string) (config MongoConfig) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	var mConfig MongoConfig
	err = json.Unmarshal(file, &mConfig)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}

	return mConfig
}

func LoadSecurityConfig(path string) (authConfig AuthConfig) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	var config AuthConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}

	//signKey = []byte(config.SigningKey)
	//return signKey
	return config
}

func LoadServerConfig(path string) (serverConfig ServerConfig) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
		fmt.Println("Config File Missing. ", err)
	}

	var config ServerConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
		fmt.Println("Config Parse Error: ", err)
	}

	return config
}

func Configuration() (c Config) {
	return config
}
