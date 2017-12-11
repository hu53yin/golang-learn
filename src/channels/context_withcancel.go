package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

//https://play.golang.org/p/wvGmvMzIMW