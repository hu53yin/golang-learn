package main

import "fmt"

func main() {
	/*if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	if 2 != 2 {
		fmt.Println("006")
	}

	if !(2 == 2) {
		fmt.Println("007")
	}

	if !(2 != 2) {
		fmt.Println("008")
	}*/

	/*
		if x := 42; x == 42 {
			fmt.Println("001")
		}
	*/

	/*x := 40
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Our value was 41")
	} else {
		fmt.Println("Our value was not 40")
	}*/

	/*for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}*/

	/*case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough*/

	/*n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}*/

	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}

//https://play.golang.org/p/8DqYVZPWDd