package main

import "fmt"

func main() {
	// Считываем целое число от пользователя - размер стороны доски
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Printf("Input error: %s\n", err)
		return
	}

	str := ""
	// проходим по столбцам
	for i := 0; i < size; i++ {
		// проходим по строке
		for j := 0; j < size; j++ {
			// записываем в строку соотвествующий символ
			if j%2 != 0 {
				str += "#"
			}
			str += " "
		}
		// переходим на новую строку и если она четная, то начинаем ее с пробела
		str += "\n"
		if i%2 == 0 {
			str += " "
		}
	}

	fmt.Println(str)
}
