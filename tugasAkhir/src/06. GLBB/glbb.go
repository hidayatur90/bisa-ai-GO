package main

import "fmt"

func findGLBB(v0 float64, a float64, t float64) float64 {
	v := v0 + (a * t)
	return v
}

func main() {
	var kecAwal float64
	var percepatan float64
	var waktu float64

	fmt.Print("Kecepatan awal (m/s) : ")
	fmt.Scanln(&kecAwal)
	fmt.Print("Percepatan (m/s^2) : ")
	fmt.Scanln(&percepatan)
	fmt.Print("Waktu (s) : ")
	fmt.Scanln(&waktu)

	glbb := findGLBB(kecAwal, percepatan, waktu)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Kecepatan setelah t = ", waktu, " adalah ", glbb, "m/s")
}
