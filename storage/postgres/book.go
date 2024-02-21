package postgres

import (
	"context"
	"develop/api/models"
	"develop/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool)storage.IBookStorage{
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) Create(ctx context.Context, createBook models.CreateBook) (string, error)  {
	query := `INSERT INTO books (id, name, author_name, page_number)
		values ($1, $2, $3, $4)
	`
	uid := uuid.New()

	_, err := b.db.Exec(ctx, query,
		uid,
		createBook.Name,
		createBook.AuthorName,
		createBook.PageNumber,
	)

	if err != nil{
		fmt.Println("error while creating book!", err.Error())
		return "",err
	}

	return uid.String(), nil
}

func (b *bookRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.Book, error) {
	book := models.Book{}

	query := `SELECT id, name, author_name, page_number, created_at, updated_at
	 from books 
			where id = $1 AND deleted_at = 0 `
	err := b.db.QueryRow(ctx,query, pKey.ID).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorName,
		&book.PageNumber,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil{
		fmt.Println("error while getting book by id!",err.Error())
		return models.Book{},err
	}

	return book, nil
}

func (b *bookRepo) GetList(ctx context.Context, req models.GetListRequest) (models.BooksResponse, error) {
	var(
		count int
		limit = req.Limit
		offset = (req.Page-1) * req.Limit
		books = []models.Book{}
	)

	// SELECTING COUNT

	countQuery := `SELECT count (1) from books where deleted_at = 0 `

	if req.Search != ""{
		countQuery += fmt.Sprintf(`AND (name ilike '%%%s%%' or author_name ilike '%%%s%%') `, req.Search, req.Search)
	}

	err := b.db.QueryRow(ctx, countQuery).Scan(
		&count,
	)
	if err != nil{
		fmt.Println("error while getting count of books!", err.Error())
		return models.BooksResponse{}, err
	}

	// SELECTING BOOKS
	tempQuery := `SELECT id, name, author_name, page_number, created_at, updated_at
	from books WHERE deleted_at = 0 `

	if req.Search != ""{
		tempQuery += fmt.Sprintf(`AND (name ilike '%%%s%%' OR author_name ilike '%%%s%%') `, req.Search, req.Search)
	}

	tempQuery += `LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(ctx, tempQuery, limit, offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
		return models.BooksResponse{},err
	}

	for rows.Next(){
		book := models.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.AuthorName,
			&book.PageNumber,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.BooksResponse{},err
		}

		books = append(books, book)
	}

	return models.BooksResponse{
		Books: books,
		Count: count,
	}, nil

}

func (b *bookRepo) Update(ctx context.Context, updBook models.UpdateBook)(string, error) {
	query := `UPDATE books SET name = $1, author_name = $2, updated_at = now()
			where id = $3 and deleted_at = 0 `
	_, err := b.db.Exec(ctx,query,
		updBook.Name,
		updBook.AuthorName,
		updBook.ID,
	)
	if err != nil{
		fmt.Println("error while updating books!", err.Error())
		return "",err
	}

	return updBook.ID, nil
}

func (b *bookRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (string,error) {
	query := `UPDATE books set deleted_at = 1 where id = $1 and deleted_at = 0 `
	_, err := b.db.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting books!", err.Error())
		return "",err
	}
	return pKey.ID, nil
	
}

 func (b *bookRepo) UpdatePageNumber(ctx context.Context, updPageNum models.UpdatePageNumber) (string,error) {
	query := `UPDATE books set page_number = $1 where id = $2 and deleted_at = 0`
	_, err := b.db.Exec(ctx, query,updPageNum.PageNumber, updPageNum.ID)
	if err != nil{
		fmt.Println("error while updating page number!", err.Error())
		return "", err
	}
	return updPageNum.ID, nil
 }