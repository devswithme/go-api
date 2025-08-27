package book

import "fyque/model/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	FindByIdAndUserId(id string, user_id string) (*domain.Book, error)
	FindByUserId(id string) (*domain.Book, error)
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
}