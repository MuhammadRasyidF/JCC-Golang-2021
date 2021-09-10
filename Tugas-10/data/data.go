package data

import (
	"math"
	"strconv"
	"strings"
)

// struct dan method soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Lebar * p.Panjang
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Lebar + p.Panjang)
}

func (t Tabung) Volume() float64 {
	return 3.14 * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Lebar * b.Panjang * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2*(b.Panjang+b.Tinggi) + 2*(b.Panjang+b.Lebar) + 2*(b.Lebar+b.Tinggi))
}

// struct dan method soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type DataPhone interface {
	PrintData() string
}

func (p Phone) PrintData() string {
	var typeColor = strings.Join(p.Colors, ", ")
	var tahun = strconv.Itoa(p.Year)
	var data string = "Name : " + p.Name + "\n" + "Brand : " + p.Brand + "\n" + "Year : " + tahun + "\n" + "Colors : " + typeColor
	return data
}

// func soal 3
func LuasPersegi(num int, isTrue bool) interface{} {
	if isTrue == true {
		if num == 0 {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return "luas persegi dengan sisi " + strconv.Itoa(num) + " cm adalah " + strconv.Itoa(num*num) + " cm"
		}
	} else {
		if num == 0 {
			return num
		} else {
			return num
		}
	}
}
