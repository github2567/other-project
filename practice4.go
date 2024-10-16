package main

func main() {

	// var msg string
	// fmt.Println(msg == "")
	// fmt.Printf("%q", msg)

	// lim := 225.0
	// v := math.Pow(10, 2)
	// fmt.Println("\n", v, " ", v > lim)

	// if v = math.Pow(20, 2); v > lim {
	// 	fmt.Println("v > lim, v = ", v)
	// }

	ratings := 8.4

	switch {
	case ratings < 5.0:
		print("Disappointed")
	case ratings >= 5.0 && ratings < 7.0:
		print("Normal")
	case ratings >= 7.0 && ratings < 10.0:
		print("Good")
	default:
		print("I don't know.")
	}
}
