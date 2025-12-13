package ui

import "fmt"

func Render(tick int, ants int, food int, eggs int) {
	Clear()

	fmt.Println("ANT COLONY SIM")
	fmt.Println("Tick:", tick)
	fmt.Println()

	fmt.Println("[ STATS ]")
	fmt.Printf("Ants : %d\n", ants)
	fmt.Printf("Food : %d\n", food)
	fmt.Printf("Eggs : %d\n", eggs)
	fmt.Println()

}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
