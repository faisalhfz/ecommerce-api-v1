package usecase

import (
	"ecommerce-api/src/entity"
	"ecommerce-api/src/entity/request"
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/repository"

	"github.com/jinzhu/copier"
)

type IProductUsecase interface {
	CreateProduct(productRequest request.CreateProductRequest) (response.ProductIDResponse, error)
	GetProducts(filter request.ProductFilterRequest) ([]response.GetProductResponse, error)
	GetProductById(id int) (response.GetProductDetailResponse, error)
	EditProductById(id int, productRequest request.CreateProductRequest) error
	DeleteProductById(id int) error
	AddProductToCart(productId int, quantity int) (response.CartIDResponse, error)
	IsProductInCart(productId int) bool
}

type ProductUsecase struct {
	pRepository repository.IProductRepository
	oRepository repository.IOrderRepository
	cRepository repository.ICartRepository
}

func NewProductUsecase(pRepository repository.IProductRepository, oRepository repository.IOrderRepository, cRepository repository.ICartRepository) *ProductUsecase {
	return &ProductUsecase{pRepository, oRepository, cRepository}
}

func (pUsecase ProductUsecase) CreateProduct(productRequest request.CreateProductRequest) (response.ProductIDResponse, error) {
	product := entity.Product{}
	copier.Copy(&product, &productRequest)
	id, err := pUsecase.pRepository.CreateProduct(product)
	if err != nil {
		return response.ProductIDResponse{}, err
	}
	return response.ProductIDResponse{ProductID: id}, nil
}

func (pUsecase ProductUsecase) GetProducts(filter request.ProductFilterRequest) ([]response.GetProductResponse, error) {
	products, _ := pUsecase.pRepository.GetAllProducts(filter)
	productResponse := []response.GetProductResponse{}
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (pUsecase ProductUsecase) GetProductById(id int) (response.GetProductDetailResponse, error) {
	product, err := pUsecase.pRepository.GetProduct(id)
	productResponse := response.GetProductDetailResponse{}
	copier.Copy(&productResponse, &product)
	if err != nil {
		return response.GetProductDetailResponse{}, err
	}
	return productResponse, nil
}

func (pUsecase ProductUsecase) EditProductById(id int, productRequest request.CreateProductRequest) error {
	product, _ := pUsecase.pRepository.GetProduct(id)
	productNew := entity.Product{}
	copier.Copy(&productNew, &productRequest)
	if err := pUsecase.pRepository.UpdateProduct(product, productNew); err != nil {
		return err
	}
	return nil
}

func (pUsecase ProductUsecase) DeleteProductById(id int) error {
	product, _ := pUsecase.pRepository.GetProduct(id)
	if err := pUsecase.pRepository.DeleteProduct(product); err != nil {
		return err
	}
	return nil
}

func (pUsecase ProductUsecase) AddProductToCart(productId int, quantity int) (response.NewOrderResponse, error) {
	product, _ := pUsecase.pRepository.GetProduct(productId)

	orderRequest := request.CreateOrderRequest{Product: product, Quantity: quantity}
	orderId, err := pUsecase.oRepository.CreateOrder(orderRequest)
	if err != nil {
		return response.NewOrderResponse{}, err
	}

	order, err := pUsecase.oRepository.GetOrder(orderId)
	cartId, err := pUsecase.cRepository.AddOrder(order)
	if err != nil {
		return response.NewOrderResponse{}, err
	}
	return response.NewOrderResponse{CartID: cartId, OrderID: orderId}, nil
}

func (pUsecase ProductUsecase) IsProductInCart(productId int) bool {
	cart, _ := pUsecase.cRepository.GetCart()
	for _, order := range cart.OrdersList {
		if order.Product.ID == productId {
			return true
		}
	}
	return false
}
