package usecase

import (
	"ecommerce-api/src/entity/request"
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/repository"

	"github.com/jinzhu/copier"
)

type IOrderUsecase interface {
	CreateOrder(orderRequest request.CreateOrderRequest) (int, error)
	GetOrderById(id int) (response.GetOrderResponse, error)
	DeleteOrderById(id int) error
	IsOrderInCart(productId int) bool
}

type OrderUsecase struct {
	oRepository repository.IOrderRepository
	cRepository repository.ICartRepository
}

func NewOrderUsecase(oRepository repository.IOrderRepository, cRepository repository.ICartRepository) *OrderUsecase {
	return &OrderUsecase{oRepository, cRepository}
}

func (oUsecase OrderUsecase) CreateOrder(orderRequest request.CreateOrderRequest) (int, error) {
	orderId, err := oUsecase.oRepository.CreateOrder(orderRequest)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}

func (oUsecase OrderUsecase) GetOrderById(id int) (response.GetOrderResponse, error) {
	var orderResponse response.GetOrderResponse
	order, err := oUsecase.oRepository.GetOrder(id)
	if err != nil {
		return response.GetOrderResponse{}, err
	}
	copier.Copy(&orderResponse, &order)
	return orderResponse, nil
}

func (oUsecase OrderUsecase) DeleteOrderById(id int) error {
	order, err := oUsecase.oRepository.GetOrder(id)
	if err != nil {
		return err
	}
	if err := oUsecase.oRepository.DeleteOrder(order); err != nil {
		return err
	}
	return nil
}

func (oUsecase OrderUsecase) IsOrderInCart(orderId int) bool {
	cart, _ := oUsecase.cRepository.GetCart()
	for _, order := range cart.OrdersList {
		if order.ID == orderId {
			return true
		}
	}
	return false
}
