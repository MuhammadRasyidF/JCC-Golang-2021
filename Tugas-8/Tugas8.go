package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// struct dan method soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.lebar * p.panjang
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.lebar + p.panjang)
}

func (t tabung) volume() float64 {
	return 3.14 * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.lebar * b.panjang * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2*(b.panjang+b.tinggi) + 2*(b.panjang+b.lebar) + 2*(b.lebar+b.tinggi))
}

// struct dan method soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type dataPhone interface {
	printData() string
}

func (p phone) printData() string {
	var typeColor = strings.Join(p.colors, ", ")
	var tahun = strconv.Itoa(p.year)
	var data string = "name : " + p.name + "\n" + "brand : " + p.brand + "\n" + "year : " + tahun + "\n" + "colors : " + typeColor
	return data
}

// func soal 3
func luasPersegi(num int, isTrue bool) interface{} {
	if isTrue == true {
		if num == 0 {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return "luas persegi dengan sisi 2 cm adalah 4 cm"
		}
	} else {
		if num == 0 {
			return num
		} else {
			return num
		}
	}
}

func main() {
	// soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	fmt.Println("---soal 1---")
	bangunDatar = segitigaSamaSisi{alas: 10, tinggi: 5}
	fmt.Println("=== Segitiga samasisi")
	fmt.Println("Luas     :", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())

	bangunDatar = persegiPanjang{panjang: 20, lebar: 10}
	fmt.Println("=== Persegi panjang")
	fmt.Println("Luas     :", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())

	bangunRuang = tabung{jariJari: 5, tinggi: 15}
	fmt.Println("=== Tabung")
	fmt.Println("Volume          :", bangunRuang.volume())
	fmt.Println("Luas Permukaan  :", bangunRuang.luasPermukaan())

	bangunRuang = balok{panjang: 10, lebar: 8, tinggi: 4}
	fmt.Println("=== balok")
	fmt.Println("Volume          :", bangunRuang.volume())
	fmt.Println("Luas Permukaan  :", bangunRuang.luasPermukaan())
	fmt.Println()

	// soal 2
	var newPhone dataPhone
	var warna = []string{"White", "Blue", "Grey"}

	newPhone = phone{name: "Xiaomi Poco X3", brand: "Xiaomi Poco Phone", year: 2019, colors: warna}
	fmt.Println("---soal 2---")
	fmt.Println(newPhone.printData())
	fmt.Println()

	// soal 3
	fmt.Println("---soal 3---")
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))
	fmt.Println()

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var angkaPertama = strconv.Itoa(kumpulanAngkaPertama.([]int)[0]) + " + " + strconv.Itoa(kumpulanAngkaPertama.([]int)[1])

	var angkaKedua = strconv.Itoa(kumpulanAngkaKedua.([]int)[0]) + " + " + strconv.Itoa(kumpulanAngkaKedua.([]int)[1])

	var jumlah int = kumpulanAngkaPertama.([]int)[0] + kumpulanAngkaPertama.([]int)[1] + kumpulanAngkaKedua.([]int)[0] + kumpulanAngkaKedua.([]int)[1]

	fmt.Println("---soal 4---")
	fmt.Println(prefix.(string)+angkaPertama+" + "+angkaKedua, "=", jumlah)
	fmt.Println()
}
