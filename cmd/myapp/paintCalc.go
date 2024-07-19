// игра - в которой необходимо угадать число за 10 попыток
package main

import (
	"fmt"
)

func paintNeeded(width, height float64) float64 {

	area := width * height
	return area / 10
}

func main() {

	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f liters\n", total)

}
