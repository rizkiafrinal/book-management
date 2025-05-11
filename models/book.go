package models

type Book struct {
    ID          int
    Title       string
    Description string
    ImageURL    string
    ReleaseYear int
    Price       int
    TotalPage   int
    Thickness   string
    CategoryID  int
    CreatedAt   string
    CreatedBy   string
    ModifiedAt  string
    ModifiedBy  string
}