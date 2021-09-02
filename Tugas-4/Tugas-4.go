package main

import (
	"fmt"
)

func main() {
	// soal 1

	// A. Jika angka ganjil maka tampilkan JCC
	// B. Jika angka genap maka tampilkan Candradimuka
	// C. Jika angka yang sedang ditampilkan adalah kelipatan 3 DAN angka ganjil maka tampilkan I Love Coding.

	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - JCC\n", i)
		} else {
			fmt.Printf("%d - Candradimuka\n", i)
		}
	}
	fmt.Println() //enter untuk misahin antar soal pada output

	// soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}
	fmt.Println() //enter untuk misahin antar soal pada output

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var i = 2
	var j = 0
	var newKalimat = make([]string, len(kalimat)-2)
	for i < len(kalimat) {
		newKalimat[j] = kalimat[i]
		i++
		j++
	}
	fmt.Println(newKalimat)
	fmt.Println() //enter untuk misahin antar soal pada output

	// soal 4
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d. %s\n", i+1, sayuran[i])
	}
	fmt.Println() //enter untuk misahin antar soal pada output

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	satuan["Volume Balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	for key, val := range satuan {
		fmt.Println(key, " = ", val)
	}
	fmt.Println() //enter untuk misahin antar soal pada output
}
