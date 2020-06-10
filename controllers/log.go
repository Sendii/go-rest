package controllers

import (
	"os"
	"log"
)

func CreateLog(){

	logpath := "./log/"
	os.MkdirAll(logpath, 0777)

	file, err := os.OpenFile("./log/info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Print("Logging to a file by ", hostname)
}