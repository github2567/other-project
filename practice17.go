package main17

import (
	"fmt"
)

func skills(xs ...string) {
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	for _, s := range xs {
		fmt.Println("skill:", s)
	}
}

func main() {

	// const (
	// 	Sunday = iota + 1
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )

	const (
		_ = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	defer fmt.Println("hello 3")

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	skills("js", "go", "python")

}
