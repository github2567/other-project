package main

import (
	"fmt"
	"strconv"
)

func main() {

	v := "72"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\n", s)

	n := strconv.Itoa(s)

	fmt.Printf("%s\n", n)
}
