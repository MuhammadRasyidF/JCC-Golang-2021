package models

type (
	NilaiMahasiswa struct {
		Nama        string `json:"Nama"`
		MataKuliah  string `json:"MataKuliah"`
		IndeksNilai string `json:"IndeksNilai"`
		Nilai       uint   `json:"Nilai"`
		ID          uint   `json:"ID"`
	}
)
