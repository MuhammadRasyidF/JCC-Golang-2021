package movie

import (
	"CRUD-REST-API/config"
	"CRUD-REST-API/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "data"
	layoutDateTime = "2006-01-02 15:04:05"
)

var specialID uint = 125

// GetAll
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {
	var movies []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Nama DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		specialID += 12
		var movie models.NilaiMahasiswa
		if err = rowQuery.Scan(
			&movie.ID,
			&movie.MataKuliah,
			&movie.Nama,
			&movie.Nilai); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		movie.ID = uint(specialID)

		movies = append(movies, movie)
	}
	return movies, nil
}

// Insert Movie
func Insert(ctx context.Context, movie models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (Nama, MataKuliah, ID) values('%s',%s, %d)", table,
		movie.Nama,
		movie.MataKuliah,
		movie.ID)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func Update(ctx context.Context, movie models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set Nama ='%s',  MataKuliah = %s, ID = %d",
		table,
		movie.Nama,
		movie.MataKuliah,
		movie.ID)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
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
