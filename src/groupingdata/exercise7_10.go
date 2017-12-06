package main

import (
	"fmt"
)

func main() {
	/*
		xs1 := []string{"James", "Bond", "Shaken, not stirred"}
		xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

		fmt.Println(xs1)
		fmt.Println(xs2)

		xxs := [][]string{xs1, xs2}
		fmt.Println(xxs)

		for i, xs := range xxs {
			fmt.Println("record: ", i)

			for j, val := range xs {
				fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
			}
		}
	*/

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	delete(m, `no_dr`)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}

//https://play.golang.org/p/6u7JYny-Bh