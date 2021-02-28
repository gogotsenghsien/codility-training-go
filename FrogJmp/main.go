package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	dist := Y - X
	if dist%D == 0 {
		return dist / D
	}
	return (dist / D) + 1
}
