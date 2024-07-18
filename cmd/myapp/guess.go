// игра - в которой необходимо угадать число за 10 попыток
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// генерируем число от 1 до 100
	target := rand.Intn(100) + 1
	fmt.Println("Было загадано число от 1 до 100.")
	fmt.Println("Сможете ли Вы угадать его?")
	fmt.Println(target)
	succsess := false
	// Основной цикл игры (даётся 10 попыток)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Количество оставшихся попыток", 10-guesses)
		fmt.Print("Введите Ваше число ")
		// получаем введёное от пользователя число из консоли
		reader := bufio.NewReader(os.Stdin)   // создаём объект для получения данных с консоли
		input, err := reader.ReadString('\n') // получаем введёное число
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // удаляем пробельные символы с концов
		guess, err := strconv.Atoi(input) // пытаемся перевести строку в тип int

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Ваше число меньше чем искомое ")
		} else if guess > target {
			fmt.Println("Ваше число больше чем искомое")
		} else {
			fmt.Print("Отлично вы угадали")
			succsess = true
			break
		}

	}
	if !succsess {
		fmt.Println("К сожалению Вы не угадали(")
		fmt.Print("Искомое число - ", target)
	}
}
