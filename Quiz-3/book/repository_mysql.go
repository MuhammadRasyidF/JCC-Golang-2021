package movie

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	table          = "book"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string
		var isValid = false
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		if book.Release_year < 1980 && book.Release_year > 2021 { // cek release book
			isValid = true
			fmt.Println("Minimal buku tahun 1980 sampai 2021")
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		pages, _ := strconv.Atoi(book.Total_page)

		switch { //menentukan ketebalanm buku
		case pages <= 100:
			book.Kategori_ketebalan = "tipis"
		case pages > 100 && pages <= 200:
			book.Kategori_ketebalan = "sedang"
		case pages > 200:
			book.Kategori_ketebalan = "tebal"
		}

		if !isValid { // masukkan dalam slice of book
			books = append(books, book)
		}
	}
	return books, nil
}

// Insert Book
func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, created_at, updated_at) values('%v',%v, %v, %d, %s, %v, NOW(), NOW())", table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title ='%s', description = %s, Image_url = %s, release_year = %d, Price = %s, total_page = %s, updated_at = NOW()",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page)

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
