package main5

import "fmt"

func add(x, y string) (string, int) {
	fmt.Println("Hello add", x, y)
	return x + " " + y, 999
}

func main() {
	// fmt.Println("Hello main", add("title", "text"))

	// var addvalue string
	// addvalue, _ = add("title", "text")
	// fmt.Println("Hello main", addvalue)

	// fmt.Println("Hello main", del())

	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}

func emote(ratings float32) string {
	switch {
	case ratings < 5.0:
		return "Disappointed"
	case ratings >= 5.0 && ratings < 7.0:
		return "Normal"
	case ratings >= 7.0 && ratings < 10.0:
		return "Good"
	default:
		return "I don't know."
	}
}

func del() (val int) {
	fmt.Println("Hello del")
	val = 999
	return
}
