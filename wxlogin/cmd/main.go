package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func g() string {
	i := rand.IntN(888889)
	return strconv.Itoa(i + 100000)
}

func main() {
	for j := 0; j < 10; j++ {
		fmt.Println(g())
	}
}
