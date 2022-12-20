package main

import "fmt"

func findKecepatan(s float64, t float64) float64 {
	v := s / t
	return v
}

func main() {
	var jarak float64
	var waktu float64

	fmt.Print("Jarak : ")
	fmt.Scanln(&jarak)
	fmt.Print("Waktu : ")
	fmt.Scanln(&waktu)

	glb := findKecepatan(jarak, waktu)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Kecepatan  = ", glb, "m/s")
}
