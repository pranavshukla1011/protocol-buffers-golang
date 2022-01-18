package main

import (
	"google.golang.org/protobuf/proto"
	"io/fs"
	"io/ioutil"
	"log"
	protobuffgo "protobuff-go/protobuff-go"
)

func main() {
	message := sampleFunc()
	err := writeToFile("sample.bin", message)
	if err != nil {
		log.Fatal("Error in writeToFile")
	}
	sm := &protobuffgo.Sample{}
	err = readFromFile("sample.bin", sm)
	if err != nil {
		log.Fatal("Error in readFromFile")
	}
	log.Println("sm : ", sm)
}

func writeToFile(fname string, pb proto.Message) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Can't serialize data", err)
		return err
	}
	err = ioutil.WriteFile(fname, data, fs.ModePerm)
	if err != nil {
		log.Fatal("Can't write data to file", err)
		return err
	}
	log.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, message proto.Message) error {
	bin, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("Cannot read from file", err)
		return err
	}

	err = proto.Unmarshal(bin, message)
	if err != nil {
		log.Fatal("Cannot de-searlize data", err)
		return err
	}
	return nil
}

func sampleFunc() *protobuffgo.Sample {
	sm := protobuffgo.Sample{
		Id:         1234,
		IsSimple:   true,
		Name:       "Pranav Shukla",
		SampleList: []int32{1, 2, 3, 4},
	}
	return &sm
}
