package main

import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	for {
		// if open == true, executes fmt.Println().
		if val, open := <-c; open {
			fmt.Println(val, open)
		} else {
			break
		}

	}

	fmt.Println()
	fmt.Println("=======================")

}
