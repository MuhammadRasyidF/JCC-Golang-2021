package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	var Kata1 string = "Jabar"
	var Kata2 string = "Coding"
	var Kata3 string = "Camp"
	var Kata4 string = "Golang"
	var Kata5 string = "2021"

	fmt.Println(Kata1, Kata2, Kata3, Kata4, Kata5)
	fmt.Println()

	//soal 2
	halo := "Halo Dunia"

	fmt.Println(strings.Replace(halo, "Dunia", "Golang", -1))
	fmt.Println()

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)
	fmt.Println()

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, err := strconv.Atoi(angkaPertama)
	angka2, err := strconv.Atoi(angkaKedua)
	angka3, err := strconv.Atoi(angkaKetiga)
	angka4, err := strconv.Atoi(angkaKeempat)

	if err == nil {
		fmt.Println(angka1 + angka2 + angka3 + angka4)
	}
	fmt.Println()

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", 2)

	fmt.Printf("\"%s\" - %d", kalimat, angka)
	fmt.Println()
}
