package main16

import (
	"encoding/json"
	"fmt"
)

// type Course struct {
// 	Name       string `json:"name"`
// 	level      string
// 	Views      int `json:"view"`
// 	instructor string
// 	fullPrice  int
// }

type movie struct {
	Title       string
	Year        int
	Rating      float32
	Genres      []string
	IsSuperHero bool
}

func main() {

	body := `[
	{
		"imdbID":       "tt4154796",
		"title":      	"Avengers: Endgame",
		"year":      	2019,
		"rating": 		8.4,
		"genres":  		["Action","Drama"],
		"isSuperHero":  true
	},
	{
		"imdbID":       "tt4154756",
		"title":      	"Avengers: Infinity War",
		"year":      	2018,
		"rating": 		8.4,
		"genres":  		["Action","Sci-Fi"],
		"isSuperHero":  true
	}
	]`

	ms := []movie{}

	err := json.Unmarshal([]byte(body), &ms)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", ms)

	// c := Course{
	// 	Name:       "basic go",
	// 	level:      "normal",
	// 	Views:      7239,
	// 	instructor: "สมชาย",
	// 	fullPrice:  9999,
	// }

	// data := []byte(`{
	// 	"name":       "basic go",
	// 	"level":      "normal",
	// 	"views":      7239,
	// 	"instructor": "สมชาย",
	// 	"fullPrice":  9999
	// }`)

	// b, err := json.Marshal(c)
	// fmt.Printf("type : %T \n", b)
	// fmt.Printf("byte : %v \n", b)
	// fmt.Printf("string : %v \n", string(b))
	// fmt.Printf("string : %s \n", b)
	// fmt.Println(err)

	// var d Course
	// err2 := json.Unmarshal(data, &d)
	// fmt.Printf("% #v\n", d)
	// fmt.Printf("% #v\n", d.Name)
	// fmt.Println(err2)
}
