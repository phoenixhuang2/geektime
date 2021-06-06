package main

import (
	// "errordemo/po"
	"errordemo/service"
	"fmt"
)

func main() {
	// var person []po.Person
	// person, err := service.SelectPerson(13)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(person)


	userName, err := service.QueryPerson(3)
	if err != nil {
		fmt.Printf("%+v", err)
		// fmt.Println(err)
		return
	}
	fmt.Println(userName)
}