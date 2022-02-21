package config

import (
	"log"
	"os"
)

func LoadAppConfigs() {
	err := Loader(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func GetServerAndPort() (string, string) {
	serverHost := GetAppProperty("SERVER_HOST")
	serverPort := GetAppProperty("SERVER_PORT")
	return serverHost, serverPort
}

func GetAppProperty(property string) string {
	return os.Getenv(property)
}
