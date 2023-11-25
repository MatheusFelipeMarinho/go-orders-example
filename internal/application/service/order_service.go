package service

import (
	"github.com/asaskevich/EventBus"
	"github.com/pkg/errors"
	"github/braip.com/internal/application/command"
	"github/braip.com/internal/application/query"
	"github/braip.com/internal/domain/entity"
)

type OrderService struct {
	createOrderCmd *command.CreateOrderCommand
	getOrderQuery  *query.GetOrderQuery
	eventBus       EventBus.Bus
}

func NewOrderService(createCmd *command.CreateOrderCommand, getQuery *query.GetOrderQuery, eventBus EventBus.Bus) *OrderService {
	return &OrderService{
		createOrderCmd: createCmd,
		getOrderQuery:  getQuery,
		eventBus:       eventBus,
	}
}

func (s *OrderService) CreateOrder(order *entity.Order) error {
	if err := s.createOrderCmd.Execute(order); err != nil {
		return errors.Wrap(err, "create order")
	}

	// dispatch event
	s.eventBus.Publish("OrderCreated", order)

	return nil
}

func (s *OrderService) ListOrder() ([]entity.Order, error) {
	return s.getOrderQuery.List()
}

func (s *OrderService) GetOrder(id uint) (*entity.Order, error) {
	return s.getOrderQuery.Execute(id)
}
