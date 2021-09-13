package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// func soal 1
func addphones(phones *[]string) {
	for i := 0; i < 7; i++ {
		var name string
		fmt.Printf("Input phone name : ")
		fmt.Scan(&name)
		*phones = append(*phones, name)
	}

}

// func soal 2
func luasLingkaran(jariJari float64) (Luas float64) {
	Luas = math.Round(math.Pi * math.Pow(jariJari, 2))
	return
}

func kelilingLingkaran(jariJari float64) (Keliling float64) {
	Keliling = math.Round(2 * math.Pi * jariJari)
	return
}

func printHasil(Hasil float64, Keterangan string) {
	if Keterangan == "Luas" {
		fmt.Println("Luas Lingkaran :", Hasil)
	} else {
		fmt.Println("Keliling Lingkaran :", Hasil)
	}
}

func main() {
	// soal 1
	fmt.Println("---soal 1---")
	var phones = []string{}
	addphones(&phones)
	sort.Strings(phones)
	for i := range phones {
		fmt.Printf("%d. %s\n", i+1, phones[i])
		time.Sleep(time.Second)
	}
	fmt.Println()

	//soal 2
	fmt.Println("---soal 2---")
	var Hasil float64
	//jari-jari 7
	Hasil = luasLingkaran(7)
	printHasil(Hasil, "Luas")
	Hasil = kelilingLingkaran(7)
	printHasil(Hasil, "Keliling")
	//jari-jari 10
	Hasil = luasLingkaran(10)
	printHasil(Hasil, "Luas")
	Hasil = kelilingLingkaran(10)
	printHasil(Hasil, "Keliling")
	//jari-jari 15
	Hasil = luasLingkaran(15)
	printHasil(Hasil, "Luas")
	Hasil = kelilingLingkaran(15)
	printHasil(Hasil, "Keliling")
	fmt.Println()

	// soal 3
	fmt.Println("---soal 3---")
	var panjang = flag.Float64("panjang", 5, "masukan panjang")
	var lebar = flag.Float64("lebar", 10, "masukan lebar")
	// 5 dan 10 dianggap sebagai default dari panjang dan lebar

	flag.Parse()
	fmt.Printf("Luas persegi panjang : %.1f\n", math.Round((*panjang)*(*lebar)))
	fmt.Printf("Keliling persegi panjang : %.1f\n", math.Round(2*((*panjang)+(*lebar))))
}
