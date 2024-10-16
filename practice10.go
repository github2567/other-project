package main10

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func (m movie) info() {
	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.rating)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t%s\n", g)
	}
}

func main() {

	ae := movie{
		title:       "Avenger: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	ae.info()

}
