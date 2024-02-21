package service

import (
	"context"
	"develop/api/models"
	"develop/storage"
	"fmt"
)

type bookService struct {
	storage storage.IStorage
}

func NewBookService(storage storage.IStorage) bookService {
	return bookService{
		storage: storage,
	}
}

func (b bookService) Create(ctx context.Context, createBook models.CreateBook) (models.Book,error) {
	id, err := b.storage.Book().Create(ctx, createBook)
	if err != nil{
		fmt.Println("Error in service, while creating book!",err.Error())
		return models.Book{},err
	}

	book, err := b.storage.Book().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil{
		fmt.Println("Error in service, while getting book after creating!",err.Error())
		return models.Book{},err
	}

	return book, nil
}

func (b bookService) GetByID(ctx context.Context, id string) (models.Book, error) {
	book, err := b.storage.Book().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil{
		fmt.Println("Error in service, while getting basket by id!", err.Error())
		return models.Book{}, err
	}

	return book, nil
}

func (b bookService) GetList(ctx context.Context, req models.GetListRequest) (models.BooksResponse, error) {
	books, err := b.storage.Book().GetList(ctx, req)
	if err != nil{
		fmt.Println("Error in service, while getting books!", err.Error())
		return	models.BooksResponse{}, err
	}
	return books, nil
}

func (b bookService) Update(ctx context.Context, updBook models.UpdateBook) (models.Book, error) {
	id, err := b.storage.Book().Update(ctx, updBook)
	if err != nil{
		fmt.Println("Error in service, while updating book!", err.Error())
		return models.Book{}, err
	}

	book, err := b.storage.Book().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil{
		fmt.Println("Error in service, while getting book after updating!", err.Error())
		return models.Book{},err
	}

	return book, nil
}

func (b bookService) Delete(ctx context.Context, id string) (error) {
	_, err :=  b.storage.Book().Delete(ctx, models.PrimaryKey{ID: id})
	if err != nil{
		fmt.Println("Error in service, while deleting book!")
		return err
	}
	return nil
}

func (b bookService) UpdatePageNumber(ctx context.Context, updPageNum models.UpdatePageNumber) (models.Book,error) {
	id, err := b.storage.Book().UpdatePageNumber(ctx,updPageNum)
	if err != nil{
		fmt.Println("Error in service, while updating page number!",err.Error())
		return models.Book{},err
	}

	book, err := b.storage.Book().GetByID(ctx,models.PrimaryKey{ID: id})
	if err != nil{
		fmt.Println("Error in service, while getting book after updating page num!",err.Error())
		return models.Book{},err
	}

	return book, nil
}