package main

import (
	"fmt"
	"strings"
)

func chessBoard(s int) string {
	str := strings.Builder{}
	// проходим по столбцам
	for i := 0; i < s; i++ {
		// проходим по строке
		for j := 0; j < s; j++ {
			// записываем в строку соотвествующий символ
			if j%2 != 0 {
				str.WriteString("#")
			}
			str.WriteString(" ")
		}
		// переходим на новую строку и если она четная, то начинаем ее с пробела
		str.WriteString("\n")
		if i%2 == 0 {
			str.WriteString(" ")
		}
	}
	return str.String()
}

func main() {
	// Считываем целое число от пользователя - размер стороны доски
	fmt.Println("Введи целое число от 1 до 200:")
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Printf("Input error: %s\n", err)
		return
	}
	fmt.Println(chessBoard(size))
}
