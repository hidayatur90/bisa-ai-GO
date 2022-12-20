package main

import "fmt"

func findLuasLingkaran(r float64) float64 {
	var phi = 3.14
	return phi * r * r
}

func main() {
	var jari float64

	fmt.Print("Panjang jari-jari (cm) : ")
	fmt.Scanln(&jari)

	luas := findLuasLingkaran(jari)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Luas Persegi  = ", luas, "cm^2")
}
