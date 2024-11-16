package main

import (
	"encoding/json"
	"go-rabbitmq/src/error"
	"go-rabbitmq/src/model"
	"go-rabbitmq/src/service"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	body := model.UserDao{
		Name:        "M. Aji Perdana",
		Age:         23,
		PhoneNumber: "085695951121",
		Address:     "Lampung",
	}

	jsonBody, err := json.Marshal(body)
	error.FailOnError(err, "Failed marshal body to json")

	service.ProducerMsg(jsonBody)
}
