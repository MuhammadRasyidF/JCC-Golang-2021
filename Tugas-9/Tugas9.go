package main

import (
	"errors"
	"fmt"
	"strconv"
)

// func no 1
func printKalimat(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

// func no 2
// jika parameter kedua bernilai true maka tampilkan kalimat (asumsi sisinya 2) "keliling segitiga sama sisinya dengan sisi 2 cm adalah 6 cm"
// jika parameter kedua bernilai false maka tampilkan hanya hasil string angka saja (misal "6")
// jika parameter pertama 0 dan parameter kedua bernilai true gunakan error dengan message "Maaf anda belum menginput sisi dari segitiga sama sisi"
// jika parameter pertama 0 dan parameter kedua bernilai false maka tampilkan error dengan message "Maaf anda belum menginput sisi dari segitiga sama sisi", beserta panic yang sudah di recover
func kelilingSegitigaSamaSisi(sisi int, Check bool) (string, error) {
	defer catch()
	var keliling = sisi * 3
	check := Check

	if check == true {
		if sisi == 0 {
			return "", errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			Sisi := strconv.Itoa(sisi)
			Keliling := strconv.Itoa(keliling)

			var kalimat = "keliling segitiga sama sisinya dengan sisi " + Sisi + " cm adalah " + Keliling + " cm"
			return kalimat, nil
		}
	} else {
		if sisi == 0 {
			err := errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
			panic(err)
		} else {
			Keliling := strconv.Itoa(keliling)
			return Keliling, nil
		}
	}
}

func catch() {
	if message := recover(); message != nil {
		fmt.Println("Terjadi error", message)
	}
}

func main() {
	// soal no 1
	defer printKalimat("Candradimuka Jabar Coding Camp", 2021)

	// soal no 2

	kalimat, err := kelilingSegitigaSamaSisi(4, true)
	if err == nil {
		fmt.Println(kalimat)
	} else {
		fmt.Println(err.Error())
	}

	kalimat, err = kelilingSegitigaSamaSisi(8, false)
	if err == nil {
		fmt.Println(kalimat)
	} else {
		fmt.Println(err.Error())
	}

	kalimat, err = kelilingSegitigaSamaSisi(0, true)
	if err == nil {
		fmt.Println(kalimat)
	} else {
		fmt.Println(err.Error())
	}

	kalimat, err = kelilingSegitigaSamaSisi(0, false)
	if err == nil {
		fmt.Println(kalimat)
	} else {
		fmt.Println(err.Error())
	}
}
