package repository

import (
    "bookstore/model"
    "database/sql"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
    rows, err := r.db.Query("SELECT id, name FROM categories")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []model.Category
    for rows.Next() {
        var c model.Category
        rows.Scan(&c.ID, &c.Name)
        categories = append(categories, c)
    }
    return categories, nil
}

func (r *CategoryRepository) Create(name string) error {
    _, err := r.db.Exec("INSERT INTO categories (name) VALUES ($1)", name)
    return err
}

func (r *CategoryRepository) GetByID(id int) (*model.Category, error) {
    row := r.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id)
    var c model.Category
    if err := row.Scan(&c.ID, &c.Name); err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *CategoryRepository) Delete(id int) error {
    result, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
    if rows, _ := result.RowsAffected(); rows == 0 {
        return sql.ErrNoRows
    }
    return err
}

func (r *CategoryRepository) GetBooksByCategory(id int) ([]model.Book, error) {
    rows, err := r.db.Query("SELECT id, title, total_page, thickness FROM books WHERE category_id = $1", id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []model.Book
    for rows.Next() {
        var b model.Book
        rows.Scan(&b.ID, &b.Title, &b.TotalPage, &b.Thickness)
        books = append(books, b)
    }
    return books, nil
}