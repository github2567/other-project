package main

import (
	"fmt"
)

type EmpItem struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

type Movie struct {
	ID int64 `json:"id"`
}

func main() {

	m := &Movie{}
	resultShow := [2]Movie{}
	m.ID = 1
	resultShow[0].ID = m.ID
	fmt.Printf("%#v\n", resultShow)
	m.ID = 2
	resultShow[1].ID = m.ID
	fmt.Printf("%#v\n", resultShow)

	// var empItem EmpItem
	// emp := make([]EmpItem, 0)

	// empItem.Id = "1"
	// empItem.Name = "FirstName1"
	// empItem.Address = "Address1"
	// emp = append(emp, empItem)

	// empItem.Id = "2"
	// empItem.Name = "FirstName2"
	// empItem.Address = "Address2"
	// emp = append(emp, empItem)

	// // jsonString, _ := json.Marshal(emp)

	// // fmt.Printf("%s", jsonString)

	// var ss []int
	// if ss == nil {
	// 	fmt.Printf("%#v", ss)
	// }
}
