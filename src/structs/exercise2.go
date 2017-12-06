package main

import (
	"fmt"
)

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Huseyin",
		lastName:  "Gurkan",
		favoriteIceCreamFlavors: []string{
			"Pistachio",
			"Vanilla",
			"Strawberry",
		},
	}

	p2 := person{
		firstName: "Berkay",
		lastName:  "Gurkan",
		favoriteIceCreamFlavors: []string{
			"Lemon",
			"Sour Cherry",
			"Strawberry",
		},
	}

	/*
		fmt.Println(p1.firstName, p1.lastName)
		for i, val := range p1.favoriteIceCreamFlavors {
			fmt.Printf("\t %v \t %v \n", i, val)
		}

		fmt.Println(p2.firstName, p2.lastName)
		for i, val := range p2.favoriteIceCreamFlavors {
			fmt.Printf("\t %v \t %v \n", i, val)
		}
	*/

	m := map[string]person{
		p1.firstName: p1,
		p2.firstName: p2,
	}

	for _, val := range m {
		fmt.Println(val.firstName, val.lastName)

		for j, flavor := range val.favoriteIceCreamFlavors {
			fmt.Printf("\t %v \t %v \n", j, flavor)
		}
	}
}

//https://play.golang.org/p/8OGM4Nh4i7