package order

import (
	"gorm.io/gorm"
	"orders/internal/domain/entity"
)

type WriteRepository struct {
	db *gorm.DB
}

func NewWriteRepository(db *gorm.DB) *WriteRepository {
	return &WriteRepository{db: db}
}

func (r *WriteRepository) CreateOrder(order *entity.Order) error {
	result := r.db.Create(order)
	return result.Error
}

func (r *WriteRepository) UpdateOrder(order *entity.Order) error {
	result := r.db.Save(order)
	return result.Error
}

func (r *WriteRepository) DeleteOrder(id int) error {
	result := r.db.Delete(&entity.Order{}, id)
	return result.Error
}
