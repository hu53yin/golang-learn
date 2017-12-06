package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: false,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(t1.doors, t1.color, t1.fourWheel)
	fmt.Println(s1.doors, s1.color, s1.luxury)
}

//https://play.golang.org/p/W65g-Y-tA4