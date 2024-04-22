package repository

import (
	"database/sql"
	"quiz3/structs"
	"time"
)

func GetAllBook(db *sql.DB) (results []structs.Book, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO books (title, description, release_year, price, total_page, thickness, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	errs := db.QueryRow(sql, book.Title, book.Description, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, time.Now(), time.Now())

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE books SET title = $1, description = $2, release_year = $3, price = $4, total_page = $5, thickness = $6, category_id = $7, updated_at = $8 WHERE id = $9"
	errs := db.QueryRow(sql, book.Title, book.Description, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, time.Now(), book.ID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM person WHERE id = $1"
	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}
