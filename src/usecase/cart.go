package usecase

import (
	"ecommerce-api/src/entity"
	"ecommerce-api/src/entity/request"
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	CreateCart(cartRequest request.CreateCartRequest) (response.CartIDResponse, error)
	GetCart() (response.GetCartResponse, error)
	// AddProductToCart(productId int) error
	CheckoutCart() (response.GetCartResponse, error)
	GetCompletedCarts() ([]response.GetCartResponse, error)
	ClearCart() error
	// RemoveProductFromCart(productId int) error
	// IsProductInCart(productId int, cart *entity.Cart) bool
}

type CartUsecase struct {
	cRepository repository.ICartRepository
	pRepository repository.IProductRepository
}

func NewCartUsecase(cRepository repository.ICartRepository, pRepository repository.IProductRepository) *CartUsecase {
	return &CartUsecase{cRepository: cRepository, pRepository: pRepository}
}

func (cUsecase CartUsecase) CreateCart() (response.CartIDResponse, error) {
	cart := request.CreateCartRequest{ProductsList: []entity.ProductEntry{}, TotalPrice: 0, IsCheckout: false}
	id, err := cUsecase.cRepository.CreateCart(cart)
	if err != nil {
		return response.CartIDResponse{}, err
	}
	return response.CartIDResponse{CartID: id}, nil
}

func (cUsecase CartUsecase) GetCart() (response.GetCartResponse, error) {
	cart, err := cUsecase.cRepository.GetCart()
	if err != nil {
		return response.GetCartResponse{}, err
	}
	productsListResponse := []response.GetProductEntryResponse{}
	for _, productEntry := range cart.ProductsList {
		product, _ := cUsecase.pRepository.GetProduct(productEntry.ProductID)
		productResponse := response.GetProductResponse{}
		copier.Copy(&productResponse, &product)
		productEntryResponse := response.GetProductEntryResponse{Product: productResponse, Quantity: productEntry.Quantity}
		productsListResponse = append(productsListResponse, productEntryResponse)
	}

	cartResponse := response.GetCartResponse{ID: cart.ID, ProductsList: productsListResponse, TotalPrice: cart.TotalPrice}
	return cartResponse, nil
}

// func (cUsecase CartUsecase) AddProductToCart(productId int, quantity int) (response.CartIDResponse, error) {
// 	product, _ := cUsecase.pRepository.GetProduct(productId)

// 	var newProduct entity.Product
// 	copier.Copy(&newProduct, &product)
// 	cartId, err := cUsecase.cRepository.AddProduct(newProduct, quantity)
// 	if err != nil {
// 		return response.CartIDResponse{}, err
// 	}
// 	return response.CartIDResponse{CartID: cartId}, nil
// }

func (cUsecase CartUsecase) CheckoutCart() (response.GetCartResponse, error) {
	cart, err := cUsecase.cRepository.CheckoutCart()
	if err != nil {
		return response.GetCartResponse{}, err
	}

	productsListResponse := []response.GetProductEntryResponse{}
	for _, productEntry := range cart.ProductsList {
		product, _ := cUsecase.pRepository.GetProduct(productEntry.ProductID)
		productResponse := response.GetProductResponse{}
		copier.Copy(&productResponse, &product)
		productEntryResponse := response.GetProductEntryResponse{Product: productResponse, Quantity: productEntry.Quantity}
		productsListResponse = append(productsListResponse, productEntryResponse)
	}

	cartResponse := response.GetCartResponse{ID: cart.ID, ProductsList: productsListResponse, TotalPrice: cart.TotalPrice}
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

		productsListResponse := []response.GetProductEntryResponse{}
		for _, productEntry := range cart.ProductsList {
			product, _ := cUsecase.pRepository.GetProduct(productEntry.ProductID)
			productResponse := response.GetProductResponse{}
			copier.Copy(&productResponse, &product)
			productEntryResponse := response.GetProductEntryResponse{Product: productResponse, Quantity: productEntry.Quantity}
			productsListResponse = append(productsListResponse, productEntryResponse)
		}

		cartResponse := response.GetCartResponse{ID: cart.ID, ProductsList: productsListResponse, TotalPrice: cart.TotalPrice}
		cartsResponse = append(cartsResponse, cartResponse)
	}
	return cartsResponse, nil
}

// func (cUsecase CartUsecase) RemoveProductFromCart(productId int) error {
// 	if err := cUsecase.cRepository.RemoveProduct(productId); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (cUsecase CartUsecase) IsProductInCart(productId int, cart *entity.Cart) bool {
// 	productList := cart.Products
// 	for _, product := range productList {
// 		if product.ID == productId {
// 			return true
// 		}
// 	}
// 	return false
// }
