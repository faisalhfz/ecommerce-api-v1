package usecase

import (
	"ecommerce-api/src/entity/request"
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	CreateCart(cartRequest request.CreateCartRequest) error
	GetCart() (response.GetCartResponse, error)
	CheckoutCart() (response.GetCartResponse, error)
	GetCompletedCarts() ([]response.GetCartResponse, error)
	ClearCart() error
	// UpdateOrderInCart(orderId int, newOrder request.CreateOrderRequest) error
	RemoveOrderFromCart(orderId int) error
}

type CartUsecase struct {
	cRepository repository.ICartRepository
	oRepository repository.IOrderRepository
	pRepository repository.IProductRepository
}

func NewCartUsecase(cRepository repository.ICartRepository, oRepository repository.IOrderRepository, pRepository repository.IProductRepository) *CartUsecase {
	return &CartUsecase{cRepository, oRepository, pRepository}
}

func (cUsecase CartUsecase) CreateCart() error {
	cart := request.CreateCartRequest{TotalPrice: 0, IsCheckout: false}
	_, err := cUsecase.cRepository.CreateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (cUsecase CartUsecase) GetCart() (response.GetCartResponse, error) {
	cart, err := cUsecase.cRepository.GetCart()
	if err != nil {
		return response.GetCartResponse{}, err
	}

	ordersListResponse := []*response.GetOrderResponse{}
	for _, order := range cart.OrdersList {
		product, _ := cUsecase.pRepository.GetProduct(order.ProductID)
		productResponse := response.GetProductResponse{}
		copier.Copy(&productResponse, &product)
		orderResponse := response.GetOrderResponse{ID: order.ID, Product: &productResponse, Quantity: order.Quantity}
		ordersListResponse = append(ordersListResponse, &orderResponse)
	}
	cartResponse := response.GetCartResponse{ID: cart.ID, OrdersList: ordersListResponse, TotalPrice: cart.TotalPrice}

	return cartResponse, nil
}

func (cUsecase CartUsecase) CheckoutCart() (response.GetCartResponse, error) {
	cart, err := cUsecase.cRepository.CheckoutCart()
	if err != nil {
		return response.GetCartResponse{}, err
	}

	ordersListResponse := []*response.GetOrderResponse{}
	for _, order := range cart.OrdersList {
		product, _ := cUsecase.pRepository.GetProduct(order.ProductID)
		productResponse := response.GetProductResponse{}
		copier.Copy(&productResponse, &product)
		orderResponse := response.GetOrderResponse{ID: order.ID, Product: &productResponse, Quantity: order.Quantity}
		ordersListResponse = append(ordersListResponse, &orderResponse)
	}
	cartResponse := response.GetCartResponse{ID: cart.ID, OrdersList: ordersListResponse, TotalPrice: cart.TotalPrice}

	return cartResponse, nil
}

func (cUsecase CartUsecase) ClearCart() error {
	cart, err := cUsecase.cRepository.GetCart()
	if err != nil {
		return err
	}
	if err := cUsecase.cRepository.DeleteCart(cart); err != nil {
		return err
	}
	return nil
}

func (cUsecase CartUsecase) GetCompletedCarts() ([]response.GetCartResponse, error) {
	carts, err := cUsecase.cRepository.GetCompletedCarts()
	if err != nil {
		return nil, err
	}

	cartsResponse := []response.GetCartResponse{}

	for _, cart := range carts {
		ordersListResponse := []*response.GetOrderResponse{}

		for _, order := range cart.OrdersList {
			product, _ := cUsecase.pRepository.GetProduct(order.ProductID)
			productResponse := response.GetProductResponse{}
			copier.Copy(&productResponse, &product)
			orderResponse := response.GetOrderResponse{ID: order.ID, Product: &productResponse, Quantity: order.Quantity}
			ordersListResponse = append(ordersListResponse, &orderResponse)
		}
		cartResponse := response.GetCartResponse{ID: cart.ID, OrdersList: ordersListResponse, TotalPrice: cart.TotalPrice}

		cartsResponse = append(cartsResponse, cartResponse)
	}
	return cartsResponse, nil
}

func (cUsecase CartUsecase) RemoveOrderFromCart(orderId int) error {
	order, err := cUsecase.oRepository.GetOrder(orderId)
	if err != nil {
		return err
	}
	if err := cUsecase.cRepository.RemoveOrder(order); err != nil {
		return err
	}
	if err := cUsecase.oRepository.DeleteOrder(order); err != nil {
		return err
	}
	return nil
}
