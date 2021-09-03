package main

import (
	"fmt"
)

// func soal 1
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// func soal 2
func introduce(nama, jenis, pekerjaan, umur string) (kalimat []string) {
	if jenis == "laki-laki" {
		kalimat = append(kalimat, "Pak", nama, "adalah", "seorang", pekerjaan, "yang berusia", umur, "tahun")
	} else {
		kalimat = append(kalimat, "Bu", nama, "adalah", "seorang", pekerjaan, "yang berusia", umur, "tahun")
	}
	return
}

// func soal 3
func buahFavorit(nama string, buah ...string) (kalimat []string) {
	kalimat = append(kalimat, "halo nama saya", nama, "dan buah favorit saya adalah")
	for i := 0; i < len(buah); i++ {
		if i < len(buah)-1 {
			kalimat = append(kalimat, "\""+buah[i]+"\""+",")
		} else {
			kalimat = append(kalimat, "\""+buah[i]+"\"")
		}
	}
	return
}

func main() {
	// soal 1
	fmt.Println("--- Soal 1 ---")

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	fmt.Println()

	// Soal 2
	fmt.Println("--- Soal 2 ---")

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Println()

	// soal 3
	fmt.Println("--- Soal 3 ---")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	fmt.Println()

	// soal 4
	fmt.Println("--- Soal 4 ---")

	var dataFilm = []map[string]string{}
	// var i = 0
	var tambahDataFilm = func(title, jam, genre, tahun string) {
		var data = map[string]string{}
		data["genre"] = genre
		data["jam"] = jam
		data["tahun"] = tahun
		data["title"] = title
		dataFilm = append(dataFilm, data)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
