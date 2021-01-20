package main

import (
	"fmt"
)

func main() {

	//get burger type
	//fmt.Println(getBurger(2))

	//get menu items by type
	menuSelection := getMenuItems("")
	for _, menuItem := range menuSelection {
		fmt.Println(menuItem)
	}
}
