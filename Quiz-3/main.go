package main

import (
	"Quiz-3/models"
	"Quiz-3/movie"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/book", Auth(GetBook))
	router.POST("/book/create", PostBook)
	router.PUT("/book/:id/update", UpdateBook)
	router.DELETE("/book/:id/delete", DeleteBook)
	router.GET("/bangun-datar/:loc", BangunDatar)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// Read
// GetBook
func GetBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	movies, err := movie.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, movies, http.StatusOK)
}

// Create
// PostBook
func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Movie
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := movie.Insert(ctx, mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Update
// UpdateBook
func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := movie.Update(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteBook
func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMovie = ps.ByName("id")
	if err := movie.Delete(ctx, idMovie); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func SegitigaSamaSisi(alas int, tinggi int, hitung string, result chan int) {
	var luas = alas * tinggi / 2
	var keliling = alas * 3

	if hitung == "luas" {
		result <- luas
	} else {
		result <- keliling
	}
}

func Persegi(sisi int, hitung string, result chan int) {
	var luas = sisi * 2
	var keliling = sisi * 4

	if hitung == "luas" {
		result <- luas
	} else {
		result <- keliling
	}
}

func PersegiPanjang(panjang int, lebar int, hitung string, result chan int) {
	var luas = panjang * lebar
	var keliling = 2 * (panjang + lebar)

	if hitung == "luas" {
		result <- luas
	} else {
		result <- keliling
	}
}

func Lingkaran(jariJari int, hitung string, result chan float64) {
	var luas = 3.14 * float64(jariJari) * 2
	var keliling = 2 * 3.14 * float64(jariJari)

	if hitung == "luas" {
		result <- luas
	} else {
		result <- keliling
	}
}

func JajarGenjang(sisi, alas, tinggi int, hitung string, result chan int) {
	var luas = alas * tinggi
	var keliling = (2 * alas) + (2 * sisi)

	if hitung == "luas" {
		result <- luas
	} else {
		result <- keliling
	}
}

func BangunDatar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(chan int)
	resfloat := make(chan float64)
	loc := ps.ByName("loc")
	takeQuery := r.URL.Query()
	switch loc {
	case "segitiga-sama-sisi":
		var alas = takeQuery.Get("alas")
		var tinggi = takeQuery.Get("tinggi")
		var hitung = takeQuery.Get("hitung")

		alasInt, _ := strconv.Atoi(alas)
		tinggiInt, _ := strconv.Atoi(tinggi)

		go SegitigaSamaSisi(alasInt, tinggiInt, hitung, result)
		HasilKali := <-result
		fmt.Fprintln(w, HasilKali)
	case "persegi":
		var sisi = takeQuery.Get("sisi")
		var hitung = takeQuery.Get("hitung")

		sisiInt, _ := strconv.Atoi(sisi)

		go Persegi(sisiInt, hitung, result)
		HasilKali := <-result
		fmt.Fprintln(w, HasilKali)
	case "persegi-panjang":
		var panjang = takeQuery.Get("panjang")
		var lebar = takeQuery.Get("lebar")
		var hitung = takeQuery.Get("hitung")

		panjangInt, _ := strconv.Atoi(panjang)
		lebarInt, _ := strconv.Atoi(lebar)

		go PersegiPanjang(panjangInt, lebarInt, hitung, result)
		HasilKali := <-result
		fmt.Fprintln(w, HasilKali)
	case "lingkaran":
		var jariJari = takeQuery.Get("jariJari")
		var hitung = takeQuery.Get("hitung")

		jariJariInt, _ := strconv.Atoi(jariJari)

		go Lingkaran(jariJariInt, hitung, resfloat)
		HasilKali := <-resfloat
		fmt.Fprintln(w, HasilKali)
	case "jajar-genjang":
		var sisi = takeQuery.Get("sisi")
		var alas = takeQuery.Get("alas")
		var tinggi = takeQuery.Get("tinggi")
		var hitung = takeQuery.Get("hitung")

		sisiInt, _ := strconv.Atoi(sisi)
		alasInt, _ := strconv.Atoi(alas)
		tinggiInt, _ := strconv.Atoi(tinggi)

		go JajarGenjang(sisiInt, alasInt, tinggiInt, hitung, result)
		HasilKali := <-result
		fmt.Fprintln(w, HasilKali)
	}

}

// middleware
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "password" {
			next(w, r, ps)
			return
		}
		if uname == "editor" && pwd == "secret" {
			next(w, r, ps)
			return
		}
		if uname == "trainer" && pwd == "rahasia" {
			next(w, r, ps)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
	}
}
