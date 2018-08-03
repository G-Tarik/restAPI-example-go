package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"rest-example/myapp"
)

type Config struct {
	ApiPort string `json:"port"`
	DbURI   string `json:"dburi"`
}

const CONFIGFILE = "config.json"

func ReadConfig() *Config {
	var c Config
	file, err := os.Open(CONFIGFILE)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func main() {
	var config = ReadConfig()
	myapp.ConnectDB(config.DbURI)
	defer myapp.DisconnectDB()

	router := myapp.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Panicln(http.ListenAndServe(":"+config.ApiPort, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
