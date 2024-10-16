package main

import "fmt"

var year = 2024

func main() {

	year = 2019
	movie_name := "Avengers: Endgame"
	rating := 8.4
	thaidesc := "ซุปเปอร์ฮีไร่: "
	var typename string

	fmt.Println(
		"เรื่อง: ", movie_name, "\n",
		"ปี: ", year,
	)

	fmt.Println(
		"เรตติ้ง: ", rating,
	)

	typename = "Sci-Fi"
	fmt.Println("ประเภท: ", typename)

	fmt.Println(thaidesc, true)
}
