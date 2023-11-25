package query

import (
	"orders/internal/domain/entity"
	"orders/internal/domain/repository/order"
)

type GetOrderQuery struct {
	Repo order.ReadRepository
}

func NewGetOrderQuery(repo order.ReadRepository) *GetOrderQuery {
	return &GetOrderQuery{Repo: repo}
}

func (q *GetOrderQuery) List() ([]entity.Order, error) {
	return q.Repo.ListOrders()
}

func (q *GetOrderQuery) Execute(id uint) (*entity.Order, error) {
	return q.Repo.GetOrderById(int(id))
}
