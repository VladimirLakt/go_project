// расчет необходимой краски
package main

import (
	"fmt"
	"go_project/cmd/myapp/greeting"
	"go_project/cmd/myapp/keyboard"
	"log"
)

func main() {
	greeting.Hello()
	fmt.Print("Введите температуру в Фаренгейтах ")
	fahrenheigt, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheigt - 32) * 5 / 9
	fmt.Printf("%.2f градуса по Цельсию", celsius)

}
