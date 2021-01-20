package main

type burgers []string

func burgerTypes() burgers {
	burgerTypes := []string{"Slim Jim", "Average Joe", "Big Devon"}

	return burgerTypes
}

func getBurger(menuSelection int) string {
	burger := ""
	switch menuSelection {
	case 1:
		burger = "Slim Jim"
	case 2:
		burger = "Average Joe"
	case 3:
		burger = "Big Devon"
	default:
		burger = "Big Devon"
	}
	return burger
}
