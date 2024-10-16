package main

import (
	"fmt"
)

type EmpItem struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

func main() {
	var empItem EmpItem
	emp := make([]EmpItem, 0)

	empItem.Id = "1"
	empItem.Name = "FirstName1"
	empItem.Address = "Address1"
	emp = append(emp, empItem)

	empItem.Id = "2"
	empItem.Name = "FirstName2"
	empItem.Address = "Address2"
	emp = append(emp, empItem)

	// jsonString, _ := json.Marshal(emp)

	// fmt.Printf("%s", jsonString)

	var ss []int
	if ss == nil {
		fmt.Printf("%#v", ss)
	}
}
