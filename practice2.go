package main

import "fmt"

func main() {
	// var charrune rune = 'ก'
	// fmt.Println(charrune)

	// fmt.Printf("%c\n", charrune)

	// charrune = '😀'
	// fmt.Println(charrune)
	// fmt.Printf("%c\n", charrune)

	year := 2019
	movie_name := "Avengers: Endgame"
	rating := 8.4
	thaidesc := "ซุปเปอร์ฮีไร่: "
	var typename string

	fmt.Printf(
		"เรื่อง: %s\nปี: %d\n",
		movie_name,
		year,
	)

	fmt.Printf(
		"เรตติ้ง: %.1f\n", rating,
	)

	typename = "Sci-Fi"
	fmt.Printf("ประเภท: %s\n", typename)

	fmt.Printf(thaidesc+"%t\n", true)

	var charrune rune = '😀'
	charrune = '⭐'
	fmt.Printf("เรื่องโปรด: %c\n", charrune)
}
