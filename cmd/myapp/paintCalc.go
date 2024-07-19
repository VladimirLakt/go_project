// расчет необходимой краски
package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("Длина или ширина не могут быть меньше нуля ")
	}
	area := width * height
	return area / 10, nil
}

func main() {

	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount, err = paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f liters\n", total)
}
