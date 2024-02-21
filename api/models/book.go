package models

import (
	"time"
)

type Book struct {
	ID 		      string`json:"id"`
	Name          string`json:"name"`
	AuthorName    string`json:"author_name"`
	PageNumber       int`json:"page_number"`
	CreatedAt  time.Time`json:"created_at"`
	UpdatedAt  time.Time`json:"updated_at"`
}

type CreateBook struct {
	Name       string`json:"name"`
	AuthorName string`json:"author_name"`
	PageNumber    int`json:"page_number"`
}

type UpdateBook struct {
	ID 		   string`json:"id"`
	Name       string`json:"name"`
	AuthorName string`json:"author_name"`
}

type PrimaryKey struct {
	ID string`json:"id"`
}

type BooksResponse struct {
	Books []Book`json:"books"`
	Count    int`json:"count"`
}

type GetListRequest struct {
	Page      int`json:"page"`
	Limit     int`json:"limit"`
	Search string`json:"search"`
}

type UpdatePageNumber struct {
	ID         string`json:"id"`
	PageNumber    int`json:"page_number"`
}