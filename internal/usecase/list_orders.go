package usecase

import (
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/internal/entity"
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/pkg/events"
)

type orderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO []orderOutput

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {

	orders, err := l.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var dto ListOrdersOutputDTO

	for _, o := range orders {
		dto = append(dto, orderOutput{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return dto, nil
}
