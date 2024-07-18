// pass_fail сообщает сдал ли пользователь экзамен
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a grade: ")          // Запросить у пользователя значение
	reader := bufio.NewReader(os.Stdin)   // Создать «буферизованное средство "чтения" для получения текста с клавиатуры.
	input, err := reader.ReadString('\n') // Возвращает текст, введенный пользователем до нажатия клавиши Enter
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade <= 60 {
		status = "passing"
	} else {
		status = "false"
	}

	fmt.Println(status)
}
