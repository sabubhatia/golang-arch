package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person {
		First: "Pheng",
	}
	
	p2 := person {
		First: "Sabu",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		// We are panicing cause this is a proogrammer error.
		// If you are tryign to marshall something that cant be marshalled then this is a programmer error.
		log.Panic(err)
	}

	fmt.Println(string(bs))
}