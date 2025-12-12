package ui

import "fmt"

func DisplayStats(AmountOfAnts int, AmountOfFood int, Tick int) {
	fmt.Println("----- ", Tick, " -----")
	fmt.Println("Amount of ants :", AmountOfAnts)
	fmt.Println("Amount of Food :", AmountOfFood)
	fmt.Println("------------------------")
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
