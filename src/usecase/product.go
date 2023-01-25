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
}

type ProductUsecase struct {
	pRepository repository.IProductRepository
	cRepository repository.ICartRepository
}

func NewProductUsecase(pRepository repository.IProductRepository, cRepository repository.ICartRepository) *ProductUsecase {
	return &ProductUsecase{pRepository, cRepository}
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

func (pUsecase ProductUsecase) AddProductToCart(productId int, quantity int) (response.CartIDResponse, error) {
	product, _ := pUsecase.pRepository.GetProduct(productId)

	productEntryRequest := request.CreateProductEntryRequest{Product: *product, Quantity: quantity}
	productEntry, err := pUsecase.pRepository.CreateProductEntry(productEntryRequest)
	if err != nil {
		return response.CartIDResponse{}, err
	}

	cartId, err := pUsecase.cRepository.AddProduct(productEntry)
	if err != nil {
		return response.CartIDResponse{}, err
	}
	return response.CartIDResponse{CartID: cartId}, nil
}
