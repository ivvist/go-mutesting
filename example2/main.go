package main

import (
	"fmt"
)

func dbl(number int) int {
	return 2 * number
}

func main() {
	fmt.Println("dbl(2)=", dbl(4))
}
