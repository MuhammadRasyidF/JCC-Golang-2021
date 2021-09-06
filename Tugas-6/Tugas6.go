package main

import (
	"fmt"
	"math"
)

//func soal 1
func Luas(luasLigkaran *float64, jariJari float64) {
	*luasLigkaran = 3.14 * math.Pow(jariJari, 2)
}

func Keliling(kelilingLingkaran *float64, jariJari float64) {
	*kelilingLingkaran = 2 * 3.14 * jariJari
}

// func soal 2
func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
	} else {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
	}
}

// func soal 3
func addFruits(buah *[]string, newFruit ...string) {
	*buah = append(*buah, newFruit...)
}

// func soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	var data = map[string]string{}

	data["genre"] = genre
	data["duration"] = duration
	data["year"] = year
	data["title"] = title

	*dataFilm = append(*dataFilm, data)
}

func printMap(dataFilm []map[string]string) {
	var i = 1
	for _, item := range dataFilm {
		var j = 0
		for key, val := range item {
			if j == 0 {
				fmt.Printf("%d. ", i)
			} else {
				fmt.Printf("   ")
			}
			fmt.Println(key + " : " + val)
			j++
		}
		i++
		fmt.Println()
	}
}

func main() {
	// soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	Luas(&luasLigkaran, 10)          //hitung luas lingkaran
	Keliling(&kelilingLingkaran, 20) //hitung keliliing lingkaran

	fmt.Println("---Soal 1---")
	fmt.Println(luasLigkaran)
	fmt.Printf("%.2f", kelilingLingkaran)
	fmt.Printf("\n\n")

	// soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println("---Soal 2---")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Println()

	// soal 3
	var buah = []string{}

	addFruits(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	fmt.Println("---Soal 3---")
	for no, buah := range buah {
		fmt.Printf("%d. %s\n", no+1, buah)
	}
	fmt.Println()

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	fmt.Println("---Soal 4---")
	printMap(dataFilm)
}
