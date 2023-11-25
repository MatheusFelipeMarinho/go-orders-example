package order

import (
	"github/braip.com/internal/domain/entity"
	"gorm.io/gorm"
)

type ReadRepository struct {
	db *gorm.DB
}

func NewReadRepository(db *gorm.DB) *ReadRepository {
	return &ReadRepository{db: db}
}

func (r *ReadRepository) GetOrderById(id int) (*entity.Order, error) {
	var order entity.Order
	result := r.db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *ReadRepository) ListOrders() ([]entity.Order, error) {
	var orders []entity.Order
	result := r.db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
