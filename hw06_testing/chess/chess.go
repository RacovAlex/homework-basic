package chess

import "strings"

func chessBoard(s int) string {
	str := strings.Builder{}
	// проходим по столбцам
	for i := 0; i < s; i++ {
		// проходим по строке
		for j := 0; j < s; j++ {
			// записываем в строку соотвествующий символ
			if j%2 == 0 {
				str.WriteString("#")
			} else {
				str.WriteString(" ")
			}
		}
		// переходим на новую строку и если она четная, то начинаем ее с пробела
		str.WriteString("\n")
		if i%2 == 0 && i != (s-1) {
			str.WriteString(" ")
		}
	}
	return str.String()
}
