package storage

import (
	"context"
	"develop/api/models"
)

type IStorage interface {
	Book() IBookStorage
	Close() 
}

type IBookStorage interface {
	Create(context.Context, models.CreateBook) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Book, error)
	GetList(context.Context, models.GetListRequest) (models.BooksResponse, error)
	Update(context.Context, models.UpdateBook)(string, error)
	Delete(context.Context, models.PrimaryKey)(string, error)
	UpdatePageNumber(context.Context, models.UpdatePageNumber) (string,error)
}

 