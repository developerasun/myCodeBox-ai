package main

import "fmt"

func main() {
	// Create two buffers in channel c.
	c := make(chan string, 2)

	c <- "Hello"
	c <- "world"

	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
