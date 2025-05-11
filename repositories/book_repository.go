package repositories

import (
    "database/sql"
    "book-management/models"
)

func GetAllBooks(db *sql.DB) ([]models.Book, error) {
    rows, err := db.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var b models.Book
        if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID); err != nil {
            return nil, err
        }
        books = append(books, b)
    }
    return books, nil
}