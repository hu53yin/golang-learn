package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
		return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON %v", err))
	}
	return bs, nil
}

//https://play.golang.org/p/Vo5kIoR-sG