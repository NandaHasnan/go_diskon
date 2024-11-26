package main

import (
	promo "diskon/diskon"
	"fmt"
)

func main() {

	var totalBelanja float64

	fmt.Print("Masukkan total belanja: ")
	fmt.Scanln(&totalBelanja)

	promo.Dis(totalBelanja)
}
