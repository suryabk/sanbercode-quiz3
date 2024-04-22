package repository

import (
	"database/sql"
	"quiz3/structs"
	"time"
)

func FilterCategoriesBook(db *sql.DB, category structs.Category) (results []structs.Book, err error) {
	sql := "SELECT * FROM books WHERE category_id = $1;"

	rows, err := db.Query(sql, category.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func GetAllCategories(db *sql.DB) (results []structs.Category, err error) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO categories (name, created_at, updated_at) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, category.Name, time.Now(), time.Now())

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3"
	errs := db.QueryRow(sql, category.Name, time.Now(), category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"
	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
