package repository

import (
    "bookstore/model"
    "database/sql"
)

type BookRepository struct {
    db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]model.Book, error) {
    rows, err := r.db.Query("SELECT id, title, release_year, total_page FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []model.Book
    for rows.Next() {
        var b model.Book
        rows.Scan(&b.ID, &b.Title, &b.ReleaseYear, &b.TotalPage)
        books = append(books, b)
    }
    return books, nil
}

func (r *BookRepository) Create(b *model.Book) error {
    _, err := r.db.Exec("INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)",
        b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID)
    return err
}

func (r *BookRepository) GetByID(id int) (*model.Book, error) {
    row := r.db.QueryRow("SELECT id, title FROM books WHERE id = $1", id)
    var b model.Book
    if err := row.Scan(&b.ID, &b.Title); err != nil {
        return nil, err
    }
    return &b, nil
}

func (r *BookRepository) Delete(id int) error {
    result, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
    if rows, _ := result.RowsAffected(); rows == 0 {
        return sql.ErrNoRows
    }
    return err
}