package main

import (
	"fmt"
)

//Create new type called menu
//slice of stings (same behaviour of []strings)
type menu []string
type category struct {
	name string
	id   int
}
type categories []category

//receiver is usually 1-2 leters matching type
func (m menu) print() {
	for i, item := range m {
		fmt.Println(i, item)
	}
}

func getMenuCategoryByID() string {

	burgers := &category{"burgers", 1}
	categories := append(categories, burgers)

	for name, id := range menuCategories {
		if categoryId == id {
			return name
		}
	}
	return "no category found"
}

func getMenuItems(selection string) menu {
	menuItems := menu{}

	switch selection {
	case "burgers":
		for _, burger := range burgerTypes() {
			menuItems = append(menuItems, burger)
		}
	default:
		menuItems = append(menuItems, "No items found!")
	}

	return menuItems
}
