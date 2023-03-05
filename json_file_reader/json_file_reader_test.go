package json_file_reader

import (
	"fmt"
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	bytes, err := ReadJsonFile("./fixtures/data.json")

	if err != nil {
		t.Fatalf("unable to read json file: %s", err)
	}
	if len(bytes) == 0 {
		t.Fatal("unable to read data from json file")
	}
}

func TestReadJsonFileDoesNotExist(t *testing.T) {
	_, err := ReadJsonFile("./fixtures/data.jsn")

	if err == nil {
		t.Fatalf("something went wrong, file does not exist")
	}

	fmt.Println(err.Error())
}
