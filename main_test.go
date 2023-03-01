package main

import (
	//"soal_pemrograman/soal_nomer2"
	"testing"
)

var (
	soal2a []int  = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	soal2b string = "*\n**\n***\n****\n*****"
	soal2c []int  = []int{1, 2, 3, 4, 5, 6, 7, 8}
	soal2d string = "input : kasur ini rusak \nresult : Palindrome \n\n input : kasur it rusak \nresult : Bukan Palindrome"
)

func Test2a(t *testing.T) {
	t.Log("Bilangan prima 1 - 100 : ", Bilanganprima(100))
}

func Test2b(t *testing.T) {
	t.Log("Pola bintang :", Polabintang(5))
}

func Test2c(t *testing.T) {
	t.Log("Ascending :", Ascending([]int{1, 8, 5, 4, 3, 2, 6, 7}))
}

func Test2d(t *testing.T) {
	t.Log("Palindrome : ", Palindrome("kasur ini rusak"))
	t.Log("Palindrome : ", Palindrome("kasur itu rusak"))
}
