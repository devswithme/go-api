package book

import (
	"fyque/model/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (r *BookRepositoryImpl) Create(book *domain.Book) error {
	return r.db.Create(&book).Error
}
func (r *BookRepositoryImpl) FindByIdAndUserId(userId string, id string) (*domain.Book, error){
	book := new(domain.Book)

	if err := r.db.Where("user_id = ? AND id = ?", userId, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}
func (r *BookRepositoryImpl) FindByUserId(userId string) (*domain.Book, error){
	book := new(domain.Book)

	if err := r.db.Where("user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	return book, nil
}
func (r *BookRepositoryImpl) Update(book *domain.Book) error {
	return r.db.Save(&book).Error
}
func (r *BookRepositoryImpl) Delete(book *domain.Book) error {
	return r.db.Delete(&book).Error
}