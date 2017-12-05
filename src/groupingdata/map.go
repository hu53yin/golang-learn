package main

import "fmt"

func main() {
	/*
		m := map[string]int{
			"James":           32,
			"Miss Moneypenny": 27,
		}
		fmt.Println(m)

		fmt.Println(m["James"])

		fmt.Println(m["Barnabas"])

		v, ok := m["Barnabas"]
		fmt.Println(v)
		fmt.Println(ok)

		if v, ok := m["James"]; ok {
			fmt.Println("THIS IS THE IF PRINT", v)
		}
	*/

	/*
		m := map[string]int{
			"James":           32,
			"Miss Moneypenny": 27,
		}
		fmt.Println(m)

		fmt.Println(m["James"])

		fmt.Println(m["Barnabas"])

		v, ok := m["Barnabas"]
		fmt.Println(v)
		fmt.Println(ok)

		m["todd"] = 33

		if v, ok := m["Barnabas"]; ok {
			fmt.Println("THIS IS THE IF PRINT", v)
		}

		for k,v := range m {
			fmt.Println(k, v)
		}

		xi := []int{4,5,7,8,9,42}

		for i,v := range xi {
			fmt.Println(i, v)
		}
	*/

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)
}

//https://play.golang.org/p/RyYi7uTBJt