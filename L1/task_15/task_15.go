package main

import (
	"fmt"
	"strings"
)

var justString string

/*
Этот фрагмент кода может привести к утечке памяти
из-за того, что вы создаёте большую строку с помощью
createHugeString, а затем сохраняете только часть этой строки в переменную
justString. Однако, срез justString = v[:100] не освобождает память, занятую v,
и всё ещё сохраняет ссылку на оригинальную строку, что может привести к утечке памяти.
*/

func createHugeString(length int) string {
	result := ""
	for length > 0 {
		result += "$"
		length--
	}

	return result
}

func someCorrectFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {

	someCorrectFunc()
	fmt.Println(justString)
}
