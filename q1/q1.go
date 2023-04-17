package main

import "fmt"

func main() {

	CalculateDiscount(1500.00, []float64{800.00, 1200.00, 700.00, 1500.00})

}

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {

	var desc float64
	var soma float64 = 0
	for _, valor := range purchaseHistory {
		soma += valor
	}
	media := soma / float64(len(purchaseHistory))

	if media > 1000 {

		desc = currentPurchase * 0.8

	} else if soma > 1000 || 0 == len(purchaseHistory) {

		desc = currentPurchase * 0.9

	} else if currentPurchase <= 1000 {

		desc = currentPurchase * 0.95

	} else if currentPurchase <= 500 {

		desc = currentPurchase * 0.98

	} else if currentPurchase <= 0 {

		return 0, fmt.Errorf("Valor menor ou igual a zero!")

	}
	var vaordesc float64

	vaordesc = currentPurchase - desc

	fmt.Printf("Valor do desconto: %f\n", vaordesc)

	return currentPurchase, nil
}
