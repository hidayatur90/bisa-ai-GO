package main

import "fmt"

func findPeriode(t float64, n float64) float64 {
	T := t / n
	return T
}

func findFrekuensi(T float64) float64 {
	f := 1 / T
	return f
}

func main() {
	var waktu float64
	var getaran float64

	fmt.Print("Waktu (s) : ")
	fmt.Scanln(&waktu)
	fmt.Print("Banyak Getaran : ")
	fmt.Scanln(&getaran)

	periode := findPeriode(waktu, getaran)
	frekuesi := findFrekuensi(periode)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Periode Getaran  = ", periode, "s")
	fmt.Println("Frekuensi Getaran  = ", frekuesi, "Hz")
}
