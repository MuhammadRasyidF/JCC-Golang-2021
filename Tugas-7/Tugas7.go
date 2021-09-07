package main

import (
	"fmt"
)

// struct soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

// struct soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// method soal 2
func (s segitiga) luasSegitiga() {
	luas := s.alas * s.tinggi / 2
	fmt.Println("Luas Segitiga =", luas)
}

func (p persegi) luasPersegi() {
	luas := p.sisi * p.sisi
	fmt.Println("Luas Persegi =", luas)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	luas := pp.panjang * pp.lebar
	fmt.Println("Luas Persegi Panjang =", luas)
}

// struct soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// method soal 3
func (p *phone) addPhoneColor(colors ...string) {
	*&p.colors = append(*&p.colors, colors...)
}

func (p phone) printPhoneSpec(name, brand string, year int, colors ...string) {
	p.name = name
	p.brand = brand
	p.year = year
	p.addPhoneColor(colors...)

	fmt.Printf("Phone Name : %s\n", p.name)
	fmt.Printf("Phone Brand : %s\n", p.brand)
	fmt.Printf("Phone Year : %d\n", p.year)
	fmt.Println("Phone color :", p.colors)
}

// struct soal 4
type movie struct {
	title, genre   string
	duration, year int
}

// func soal 4
func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	var film movie
	film.title = title
	film.duration = duration
	film.genre = genre
	film.year = year

	*dataFilm = append(*dataFilm, film)
}

func printDataFilm(dataFilm []movie) {
	var i = 1
	for _, item := range dataFilm {
		fmt.Printf("%d. title : %s\n", i, item.title)
		fmt.Println("   duration :", item.duration)
		fmt.Println("   genre :", item.genre)
		fmt.Println("   year :", item.year)
		fmt.Println()
		i++
	}
}

func main() {
	// soal 1
	var buahNanas = buah{"Nanas", "Kuning", false, 9000}
	var buahJeruk = buah{"Jeruk", "Oranye", true, 8000}
	var buahSemangka = buah{warna: "Hijau & Merah", adaBijinya: true, nama: "Semangka", harga: 10000}
	var buahPisang buah
	buahPisang.nama = "Pisang"
	buahPisang.warna = "Kuning"
	buahPisang.harga = 5000
	buahPisang.adaBijinya = false

	fmt.Println("---soal 1---")
	fmt.Println(buahNanas)
	fmt.Println(buahJeruk)
	fmt.Println(buahSemangka)
	fmt.Println(buahPisang)
	fmt.Println()

	// soal 2
	var Segitiga = segitiga{alas: 5, tinggi: 10}
	var Persegi = persegi{sisi: 20}
	var Persegi_panjang = persegiPanjang{panjang: 30, lebar: 10}

	fmt.Println("---soal 2---")
	Segitiga.luasSegitiga()
	Persegi.luasPersegi()
	Persegi_panjang.luasPersegiPanjang()
	fmt.Println()

	// soal 3
	fmt.Println("---soal 3---")
	var Poco phone
	Poco.printPhoneSpec("Poco Phone", "Xiaomi", 2019, "White", "Blue", "Silver")
	fmt.Println()

	// soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println("---soal 4---")
	printDataFilm(dataFilm)
}
