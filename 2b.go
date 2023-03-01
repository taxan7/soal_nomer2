package main

import (
	"fmt"
)

func Polabintang(n int) string {

	var result string

	for col := 0; col <= n; col++ {
		for bar := 0; bar < col; bar++ {
			result += fmt.Sprint("*")

		}
		result += fmt.Sprintln()

	}

	return result
}
