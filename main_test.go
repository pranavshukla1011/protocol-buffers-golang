package main

import (
	"log"
	protobuffgo "protobuff-go/protobuff-go"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	message := sampleFunc()
	err := writeToFile("sample.bin", message)
	if err != nil {
		log.Fatal("Error in writeToFile")
	}
}

func TestReadFromFile(t *testing.T) {
	sm := &protobuffgo.Sample{}
	err := readFromFile("sample.bin", sm)
	if err != nil {
		log.Fatal("Error in readFromFile")
	}
	log.Println("sm : ", sm)
}

func TestGetJson(t *testing.T) {
	message := sampleFunc()
	jsonObj, err := GetJson(message)

	if err != nil {
		log.Fatal("Error in GetJson", err)
	}
	log.Println("JsonObj: \n", jsonObj)
}

func TestWriteToJson(t *testing.T) {
	message := sampleFunc()
	writeToJson("simple.json", message)
}
