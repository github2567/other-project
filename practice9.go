package main9

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func main() {

	var mvs []movie

	var item movie
	item.title = "Avenger: Endgame"
	item.year = 2019
	item.rating = 8.4
	item.genres = []string{"Action", "Drama"}
	item.isSuperHero = true

	mvs = append(mvs, item)

	item.title = "Avenger: Infinity War"
	item.year = 2018
	item.rating = 8.4
	item.genres = []string{"Action", "Sci-Fi"}
	item.isSuperHero = true

	mvs = append(mvs, item)

	for _, m := range mvs {
		fmt.Printf("%#v\n", m)
	}

}
