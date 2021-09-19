package models

type (
	NilaiMahasiswa struct {
		Nama       string `json:"Nama"`
		MataKuliah string `json:"MataKuliah"`
		Nilai      uint   `json:"Nilai"`
		ID         uint   `json:"ID"`
	}
)
