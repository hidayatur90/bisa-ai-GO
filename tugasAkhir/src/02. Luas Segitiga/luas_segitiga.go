package main

import "fmt"

func findLuasSegitiga(a float64, t float64) float64 {
	return 0.5 * a * t
}

func main() {
	var alas float64
	var tinggi float64

	fmt.Print("Alas (cm) : ")
	fmt.Scanln(&alas)
	fmt.Print("Tinggi (cm) : ")
	fmt.Scanln(&tinggi)

	luas := findLuasSegitiga(alas, tinggi)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Luas Persegi  = ", luas, "cm^2")
}
