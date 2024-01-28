package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("not a quad function")
		return
	}
	// take arguments from the user, check if they are valid ,
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	// convert them to x and why int

	// pass them to the function
	// do the same to all the functions

	QuadE(x, y)
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || ((i == y-1 && j == x-1) && y > 1 && x > 1) {
				z01.PrintRune('A')
			} else if (i == y-1 && j == 0) || (i == 0 && j == x-1) {
				z01.PrintRune('C')
			} else if (i == 0) || (i == y-1) {
				z01.PrintRune('B')
			} else if j == 0 || j == x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
