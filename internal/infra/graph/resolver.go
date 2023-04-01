package graph

import "github.com/Polidoro-root/go-expert-classes/18_clean_architecture/internal/usecase"

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}
