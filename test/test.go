package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Person struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

func main() {
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}

	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}

	for _, p := range persons {
		fmt.Printf("%d : %s\n", p.Id, p.Name)
	}
}
