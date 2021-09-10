package main

import (
	D "Tugas-10/data"
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var bangunDatar D.HitungBangunDatar
	var bangunRuang D.HitungBangunRuang

	fmt.Println("---soal 1---")
	bangunDatar = D.SegitigaSamaSisi{Alas: 10, Tinggi: 5}
	fmt.Println("=== Segitiga samasisi")
	fmt.Println("Luas     :", bangunDatar.Luas())
	fmt.Println("Keliling :", bangunDatar.Keliling())

	bangunDatar = D.PersegiPanjang{Panjang: 20, Lebar: 10}
	fmt.Println("=== Persegi Panjang")
	fmt.Println("Luas     :", bangunDatar.Luas())
	fmt.Println("Keliling :", bangunDatar.Keliling())

	bangunRuang = D.Tabung{JariJari: 5, Tinggi: 15}
	fmt.Println("=== Tabung")
	fmt.Println("Volume          :", bangunRuang.Volume())
	fmt.Println("Luas Permukaan  :", bangunRuang.LuasPermukaan())

	bangunRuang = D.Balok{Panjang: 10, Lebar: 8, Tinggi: 4}
	fmt.Println("=== balok")
	fmt.Println("Volume          :", bangunRuang.Volume())
	fmt.Println("Luas Permukaan  :", bangunRuang.LuasPermukaan())
	fmt.Println()

	// soal 2
	var newPhone D.DataPhone
	var warna = []string{"White", "Blue", "Grey"}

	newPhone = D.Phone{Name: "Xiaomi Poco X3", Brand: "Xiaomi Poco Phone", Year: 2019, Colors: warna}
	fmt.Println("---soal 2---")
	fmt.Println(newPhone.PrintData())
	fmt.Println()

	// soal 3
	fmt.Println("---soal 3---")
	fmt.Println(D.LuasPersegi(4, true))

	fmt.Println(D.LuasPersegi(8, false))

	fmt.Println(D.LuasPersegi(0, true))

	fmt.Println(D.LuasPersegi(0, false))
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
