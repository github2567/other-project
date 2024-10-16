package main

import "fmt"

func main() {

	genres := []string{"Action", "Adventure", "Fantasy"}

	fmt.Printf("before for loop: %#v\n", genres)

	for index := 0; index < len(genres); index++ {
		genres[index] = "Movie: " + genres[index]
	}

	fmt.Printf("after for loop: %#v\n", genres)

	for index, val := range genres {
		fmt.Printf("genres[%d]: %s\n", index, val)
	}

}
