package command

import (
	"orders/internal/domain/entity"
	"orders/internal/domain/repository/order"
)

type CreateOrderCommand struct {
	Repo order.WriteRepository
	//eventBus event.EventBus
}

func NewCreateOrderCommand(repo order.WriteRepository) *CreateOrderCommand {
	return &CreateOrderCommand{
		Repo: repo,
	}
}

func (c *CreateOrderCommand) Execute(order *entity.Order) error {
	return c.Repo.CreateOrder(order)

	// Publish event
}
