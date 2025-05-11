package model

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	CategoryID  int    `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	ModifiedAt  string `json:"modified_at"`
	ModifiedBy  string `json:"modified_by"`
}
