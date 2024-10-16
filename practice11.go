package main11

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float32
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float32) {
	m.votes = append(m.votes, rating)
}

func main() {

	eg := &movie{
		title:       "Avenger: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float32{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	eg.addVote(8)
	fmt.Println("votes:", eg.votes)

}
