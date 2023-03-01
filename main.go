package main

import (
	"fmt"
	//"soal_pemrograman/soal_nomer2"
)

func main() {

	var soal int
	var subsoal string
	fmt.Println("silahkan pilih nomer soal!")
	fmt.Scan(&soal)

	if soal == 2 {
		fmt.Println("silahkan pilih soal a/b/c/d!")
		fmt.Scan(&subsoal)
		if subsoal == "a" {

			fmt.Println("===== SOAL NOMER 2A =====")
			var n = 100
			tampung := Bilanganprima(n)
			fmt.Println("input :", n)
			fmt.Println("result :", tampung)
		} else if subsoal == "b" {
			fmt.Println("===== SOAL NOMER 2B =====")
			var n = 5
			tampung := Polabintang(n)
			fmt.Println("input :", n)
			fmt.Println("result : \n", tampung)
		} else if subsoal == "c" {
			fmt.Println("===== SOAL NOMER 2C =====")
			var n = []int{1, 8, 5, 4, 3, 2, 6, 7}
			tampung := Ascending(n)
			fmt.Println("input", n)
			fmt.Println("result :", tampung)
		} else if subsoal == "d" {
			fmt.Println("===== SOAL NOMER 2C =====")
			var n1 = "kasur ini rusak"
			var n2 = "kasur itu rusak"
			tampung1 := Palindrome(n1)
			tampung2 := Palindrome(n2)
			fmt.Println("input :", n1)
			fmt.Println("result :", tampung1)
			fmt.Println("input :", n2)
			fmt.Println("result :", tampung2)
		} else {

		}
	}
}
