package main

import "fmt"

func sincostan(de float64, mi float64, sa float64) {
	sin := de / mi
	cos := sa / mi
	tan := de / sa

	fmt.Println("============= HASIL ==============")
	fmt.Println("Sin  = ", sin, "cm")
	fmt.Println("Cos  = ", cos, "cm")
	fmt.Println("Tangen  = ", tan, "cm")
}

func main() {
	var depan float64
	var samping float64
	var miring float64

	fmt.Print("Panjang sisi depan (cm) : ")
	fmt.Scanln(&depan)
	fmt.Print("Panjang sisi miring (cm) : ")
	fmt.Scanln(&miring)
	fmt.Print("Panjang sisi samping (cm) : ")
	fmt.Scanln(&samping)

	sincostan(depan, miring, samping)
}
