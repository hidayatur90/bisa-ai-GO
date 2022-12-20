package main

import "fmt"

func findKinetik(m float64, v float64) float64 {
	Ek := 0.5 * m * v
	return Ek
}

func findPotensial(m float64, h float64) float64 {
	var g = 10.0
	Ep := m * g * h
	return Ep
}

func main() {
	var massa float64
	var ketinggian float64
	var kecepatan float64

	fmt.Print("Massa benda (kg) : ")
	fmt.Scanln(&massa)
	fmt.Print("Ketinggian (m) : ")
	fmt.Scanln(&ketinggian)
	fmt.Print("Kecepatan (m/s^2) : ")
	fmt.Scanln(&kecepatan)

	kinetik := findKinetik(massa, kecepatan)
	potensial := findPotensial(massa, ketinggian)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Energi Kinetik  = ", kinetik, "Joule")
	fmt.Println("Energi Potensial  = ", potensial, "Joule")
}
