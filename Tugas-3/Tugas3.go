package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	Panjang1, _ := strconv.Atoi(panjangPersegiPanjang)
	Lebar1, _ := strconv.Atoi(lebarPersegiPanjang)

	Alas2, _ := strconv.Atoi(alasSegitiga)
	Tinggi2, _ := strconv.Atoi(tinggiSegitiga)

	var kelilingPersegiPanjang int = 2 * (Panjang1 + Lebar1)
	var luasSegitiga int = Alas2 * Tinggi2 / 2

	fmt.Printf("Keliling persegi panjang = %d\n", kelilingPersegiPanjang)
	fmt.Printf("Luas segitiga = %d\n\n", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var Njohn = 'A' //default nilai agar bertipe char
	var NDoe = 'A'

	if nilaiJohn >= 80 {
		Njohn = 'A'
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		Njohn = 'B'
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		Njohn = 'C'
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		Njohn = 'D'
	} else {
		Njohn = 'E'
	}

	if nilaiDoe >= 80 {
		NDoe = 'A'
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		NDoe = 'B'
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		NDoe = 'C'
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		NDoe = 'D'
	} else {
		NDoe = 'E'
	}

	fmt.Printf("Nilai John = %c\n", Njohn)
	fmt.Printf("Nilai Doe = %c\n\n", NDoe)

	// soal 3

	var tanggal = 17
	var bulan = 8
	var tahun = 1945

	tanggal = 20
	bulan = 11
	tahun = 2001

	var namaBulan string //untuk menampung nama bulan

	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Febuari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	}

	fmt.Printf("%d %s %d\n\n", tanggal, namaBulan, tahun)

	// soal 4
	tahunLahir := 2001

	switch {
	case tahunLahir >= 1944 && tahunLahir <= 1964:
		fmt.Println("Baby boomer")
	case tahunLahir >= 1965 && tahunLahir <= 1979:
		fmt.Println("Generasi X")
	case tahunLahir >= 1980 && tahunLahir <= 1994:
		fmt.Println("Generasi Y (Millenials)")
	case tahunLahir >= 1995 && tahunLahir <= 2015:
		fmt.Println("Generasi Z")
	}
}
