package book

import (
	"fyque/model/domain"
	"fyque/model/web/book"
)

type BookService interface {
	Create(req book.CreateRequest) (*domain.Book, error)
	
}