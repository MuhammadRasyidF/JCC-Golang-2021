package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

// func soal 1
func addphones(phones *[]string) {
	var namePhone = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	*phones = append(*phones, namePhone...)
}

func showPhone(phones []string, wg *sync.WaitGroup) {
	for i := range phones {
		wg.Done()
		fmt.Printf("%d. %s\n", i+1, phones[i])
		// time.Sleep(time.Second)
	}
}

//func soal 2
func getMovies(moviesChannel chan<- string, movies ...string) {
	moviesChannel <- "List Movies :"
	for index, movie := range movies {
		moviesChannel <- strconv.Itoa(index+1) + ". " + movie
	}
	close(moviesChannel)
}

//func soal 3
func luasLingkaran(jariJari float64, luas chan float64) {
	luas <- math.Round(math.Pi * math.Pow(jariJari, 2))
}

func kelilingLingkaran(jariJari float64, Keliling chan float64) {
	Keliling <- math.Round(2 * math.Pi * jariJari)
}

func volumeTabung(jariJari, tinggi float64, volume chan float64) {
	volume <- math.Round(math.Pi * math.Pow(jariJari, 2) * tinggi)
}

//func soal 4
func luasPersegiPanjang(panjang, lebar float64, luasP chan float64) {
	luasP <- math.Round(panjang * lebar)
}

func kelilingPersegiPanjang(panjang, lebar float64, kelilingP chan float64) {
	kelilingP <- math.Round(2 * (panjang + lebar))
}

func volumeBalok(panjang, lebar, tinggi float64, volumeB chan float64) {
	volumeB <- math.Round(panjang * lebar * tinggi)
}

func main() {
	// soal 1
	fmt.Println("---soal 1---")
	var phones = []string{}
	var wg sync.WaitGroup

	addphones(&phones)
	sort.Strings(phones)
	wg.Add(len(phones))
	go showPhone(phones, &wg)
	wg.Wait()
	fmt.Println()

	// soal 2
	fmt.Println("---soal 2---")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
	fmt.Println()

	// soal 3
	// jari-jari yang di gunakan adalah 8, 14 dan 20
	// tinggi tabungnya adalah 10
	fmt.Println("---soal 3---")
	Luas := make(chan float64, 3)
	Keliling := make(chan float64, 3)
	Volume := make(chan float64, 3)
	var jariJari = []float64{8, 14, 20}

	// go luasLingkaran(8, Luas) //coding yang dulu
	// go luasLingkaran(14, Luas)
	// go luasLingkaran(20, Luas)

	// go kelilingLingkaran(8, Keliling)
	// go kelilingLingkaran(14, Keliling)
	// go kelilingLingkaran(20, Keliling)

	// go volumeTabung(8, 10, Volume)
	// go volumeTabung(14, 10, Volume)
	// go volumeTabung(20, 10, Volume)

	// fmt.Printf("Luas Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(8), <-Luas)
	// fmt.Printf("Luas Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(14), <-Luas)
	// fmt.Printf("Luas Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(20), <-Luas)
	// fmt.Println()
	// fmt.Printf("Keliling Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(8), <-Keliling)
	// fmt.Printf("Keliling Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(14), <-Keliling)
	// fmt.Printf("Keliling Lingkaran dengan jari-jari %.0f adalah %.0f\n", float64(20), <-Keliling)
	// fmt.Println()
	// fmt.Printf("Volume Tabung dengan jari-jari %.0f dan tinggi %.0f adalah %.0f\n", float64(8), float64(10), <-Volume)
	// fmt.Printf("Volume Tabung dengan jari-jari %.0f dan tinggi %.0f adalah %.0f\n", float64(14), float64(10), <-Volume)
	// fmt.Printf("Volume Tabung dengan jari-jari %.0f dan tinggi %.0f adalah %.0f\n", float64(20), float64(10), <-Volume)
	// fmt.Println()

	for _, JariJari := range jariJari { //update codinhg lebih simple
		go luasLingkaran(JariJari, Luas)
		fmt.Printf("Luas Lingkaran dengan jari-jari %.0f adalah %.0f\n", JariJari, <-Luas)

		go kelilingLingkaran(8, Keliling)
		fmt.Printf("Keliling Lingkaran dengan jari-jari %.0f adalah %.0f\n", JariJari, <-Keliling)

		go volumeTabung(JariJari, 10, Volume)
		fmt.Printf("Volume Tabung dengan jari-jari %.0f dan tinggi %.0f adalah %.0f\n\n", JariJari, float64(10), <-Volume)

	}

	//soal 4
	fmt.Println("---soal 4---")
	// buatlah perhitungan luas persegi panjang dan
	// keliling persegi panjang dan volume balok,
	// gunakan select untuk menentukan hasil dari luas,
	// keliling dan volume balok masing-masing
	LuasPersegiPanjang := make(chan float64)
	KelilingPersegiPanjang := make(chan float64)
	VolumeBalok := make(chan float64)

	go luasPersegiPanjang(4, 5, LuasPersegiPanjang)
	go kelilingPersegiPanjang(10, 8, KelilingPersegiPanjang)
	go volumeBalok(3, 4, 2, VolumeBalok)

	for i := 0; i < 3; i++ {
		select {
		case LuasP := <-LuasPersegiPanjang:
			fmt.Printf("Luas Persegi Panjang dengan panjang 4 dan lebar 5 adalah %.0f\n", LuasP)
		case KelilingP := <-KelilingPersegiPanjang:
			fmt.Printf("Keliling Persegi Panjang dengan panjang 10 dan lebar 8 adalah %.0f\n", KelilingP)
		case VolumeB := <-VolumeBalok:
			fmt.Printf("Volume Balok dengan panjang 3, lebar 4, dan tinggi 2 adalah %.0f\n", VolumeB)
		}
	}
}
