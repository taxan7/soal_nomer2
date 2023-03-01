package main

import "fmt"

func Palindrome(n string) string {
	var result string
	var LastLetter = ""

	for last := len(n) - 1; last >= 0; last-- {
		LastLetter += string(n[last])
	}
	for firts := range n {
		if n[firts] != LastLetter[firts] {
			result = fmt.Sprintln("Bukan Palindrome")
			return result
		}
		result = fmt.Sprintln("Palindrome")
	}

	return result
}
