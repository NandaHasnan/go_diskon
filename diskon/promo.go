package promo

import "fmt"

func Dis(totalBelanja float64) {
	total := totalBelanja - (totalBelanja / 100 * 50)
	if totalBelanja >= 100000 {
		fmt.Printf("Anda mendapatkan diskon 50%%! Total belanja setelah diskon: Rp %v", total)
	} else {
		fmt.Printf("Anda tidak mendapatkan diskon. Total belanja: Rp %v", totalBelanja)
	}
}
