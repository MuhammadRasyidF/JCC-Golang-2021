package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// data yang di input hanya boleh Nama, MataKuliah dan Nilai saja, untuk ID dan IndeksNilai harus diolah terlebih dahulu baru di masukkan ke tambahkan ke nilaiNilaiMahasiswa
// Nilai hanya boleh diisi maksimal dengan angka 100
// untuk mengisi IndeksNilai memiliki kondisi: Nilai >= 80 indeksnya A, Nilai >= 70 dan Nilai < 80 indeksnya B, Nilai >= 60 dan Nilai < 70, indeksnya c Nilai >= 50 dan Nilai < 60 indeksnya D, Nilai < 50 indeksnya E
// harus memasukkan dulu username dan password baru bisa melakukan POST Nilai Mahasiswa

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

var specialID uint = 125

func NilaiParaMahasiswa() []NilaiMahasiswa {
	MHS := []NilaiMahasiswa{
		{Nama: "Rasyid", MataKuliah: "Fisika", Nilai: 90, IndeksNilai: "A", ID: 123},
		{Nama: "Faza", MataKuliah: "Inggris", Nilai: 80, IndeksNilai: "A", ID: 374},
		{Nama: "Rudi", MataKuliah: "Bahasa", Nilai: 70, IndeksNilai: "B", ID: 555},
		{Nama: "Prasetya", MataKuliah: "Olahraga", Nilai: 50, IndeksNilai: "D", ID: 762},
	}
	return MHS
}

func MakeIndexNilai(N uint) string {
	switch {
	case N >= 80:
		return "A"
	case N >= 70 && N < 80:
		return "B"
	case N >= 60 && N < 70:
		return "C"
	case N >= 50 && N < 60:
		return "D"
	case N < 50:
		return "E"
	default:
		return "Tidak ada nilai"
	}
}

// middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "Rasyid" && pwd == "Golang" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

// Get Nilai Mahasiswa
func getMHS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataTotalMHS, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataTotalMHS)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// Post Nilai Mahasiswa
func PostMHS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NMahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&NMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			specialID += 12
			nama := r.PostFormValue("Name")
			mataKuliahMHS := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			nilaiMHS, _ := strconv.Atoi(getNilai)
			NMahasiswa = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliahMHS,
				Nilai:       uint(nilaiMHS),
				IndeksNilai: MakeIndexNilai(uint(nilaiMHS)),
				ID:          uint(specialID),
			}
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NMahasiswa)
		dataNMahasiswa, _ := json.Marshal(NMahasiswa)
		w.Write(dataNMahasiswa)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	nilaiNilaiMahasiswa = NilaiParaMahasiswa()
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/TotalNilaiMahasiswa", Auth(http.HandlerFunc(getMHS)))
	http.Handle("/NilaiMahasiswa", Auth(http.HandlerFunc(PostMHS)))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
