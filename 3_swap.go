package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, m := swap("a", "m")
	fmt.Println(a, m);
}
