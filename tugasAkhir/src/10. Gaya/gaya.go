package main

import "fmt"

func findGaya(m float64, a float64) float64 {
	F := m * a
	return F
}

func findUsaha(F float64, s float64) float64 {
	W := F * s
	return W
}

func findDaya(W float64, t float64) float64 {
	P := W / t
	return P
}

func findTekanan(F float64, A float64) float64 {
	p := F / A
	return p
}

func main() {
	var massa float64
	var percepatan float64
	var perpindahan float64
	var waktu float64
	var luasPenampang float64

	fmt.Print("Massa Benda (kg) : ")
	fmt.Scanln(&massa)
	fmt.Print("Percepatan Benda (m/s) : ")
	fmt.Scanln(&percepatan)
	fmt.Print("Perpindahan Benda (m) : ")
	fmt.Scanln(&perpindahan)
	fmt.Print("Waktu (s) : ")
	fmt.Scanln(&waktu)
	fmt.Print("Luas Penampang (m^2) : ")
	fmt.Scanln(&luasPenampang)

	gaya := findGaya(massa, percepatan)
	usaha := findUsaha(gaya, perpindahan)
	daya := findDaya(usaha, waktu)
	tekanan := findTekanan(gaya, luasPenampang)

	fmt.Println("============= HASIL ==============")
	fmt.Println("Gaya = ", gaya, "Newton")
	fmt.Println("Usaha = ", usaha, "Joule")
	fmt.Println("Daya = ", daya, "Watt")
	fmt.Println("Tekanan = ", tekanan, "N/m^2")
}
