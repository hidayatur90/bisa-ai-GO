package main

import "fmt"

func findSuhu(C float64) {
	K := C + 273.15
	F := C*1.8 + 32
	Re := C * 0.8

	fmt.Println("============= HASIL ==============")
	fmt.Println("Kelvin = ", K, "Kelvin")
	fmt.Println("Fahrenheit = ", F, "◦F")
	fmt.Println("Reamur = ", Re, "◦R")
}

func main() {
	var celcius float64

	fmt.Print("Nilai ◦C : ")
	fmt.Scanln(&celcius)

	findSuhu(celcius)
}
