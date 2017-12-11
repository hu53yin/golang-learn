package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:",runtime.GOARCH)
}

//https://play.golang.org/p/KAKcS-Vfl9