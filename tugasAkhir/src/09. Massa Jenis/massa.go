package main

import "fmt"

func findRo(m float64, v float64) float64 {
	ro := m / v
	return ro
}

func main() {
	var massa float64
	var volume float64

	fmt.Print("Massa Benda (kg) : ")
	fmt.Scanln(&massa)
	fmt.Print("Volume Benda (m^3) : ")
	fmt.Scanln(&volume)

	res_ro := findRo(massa, volume)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Massa Benda = ", res_ro, "kg/m^3")
}
