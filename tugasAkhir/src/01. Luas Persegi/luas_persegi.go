package main

import "fmt"

func findLuasPersegi(s float64) float64 {
	return s * s
}

func main() {
	var sisi float64

	fmt.Print("Panjang sisi (cm) : ")
	fmt.Scanln(&sisi)

	luas := findLuasPersegi(sisi)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Luas Persegi  = ", luas, "cm^2")
}
