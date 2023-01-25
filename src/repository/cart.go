package repository

import (
	"ecommerce-api/src/entity"
	"ecommerce-api/src/entity/request"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(cartRequest request.CreateCartRequest) (int, error)
	GetCart() (*entity.Cart, error)
	AddProduct(productEntry entity.ProductEntry) (int, error)
	CheckoutCart() (*entity.Cart, error)
	GetCompletedCarts() ([]entity.Cart, error)
	DeleteCart(cart *entity.Cart) error
	// IsProductExist(productId int) (bool, error)

	// RemoveProduct(productId int) error
	// GetProductsList() ([]entity.ProductEntry, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (cRepository CartRepository) CreateCart(cartRequest request.CreateCartRequest) (int, error) {
	var cart entity.Cart
	copier.Copy(&cart, &cartRequest)
	if err := cRepository.db.Debug().Create(&cart).Error; err != nil {
		return 0, err
	}
	cRepository.db.Model(&cart).Association("ProductsList").Append(&cart.ProductsList)
	return cart.ID, nil
}

func (cRepository CartRepository) GetCart() (*entity.Cart, error) {
	var cart *entity.Cart
	if err := cRepository.db.Preload("ProductsList").Where("is_checkout = ?", false).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (cRepository CartRepository) AddProduct(productEntry entity.ProductEntry) (int, error) {
	cart, _ := cRepository.GetCart()

	cart.ProductsList = append(cart.ProductsList, productEntry)
	cart.TotalPrice += productEntry.Product.Price * productEntry.Quantity

	cRepository.db.Model(&cart).Association("ProductsList").Replace(&cart.ProductsList)
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func (cRepository CartRepository) CheckoutCart() (*entity.Cart, error) {
	cart, err := cRepository.GetCart()
	if err != nil {
		return nil, err
	}
	cRepository.db.Preload("ProductsList").First(&cart)
	cart.IsCheckout = true
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (cRepository CartRepository) GetCompletedCarts() ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := cRepository.db.Preload("ProductsList").Where("is_checkout = ?", true).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (cRepository CartRepository) DeleteCart(cart *entity.Cart) error {
	var products []entity.Product
	cRepository.db.Model(entity.Product{}).Find(&products)
	cRepository.db.Model(&cart).Association("Languages").Delete(products)
	if err := cRepository.db.Debug().Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}

// func (cRepository CartRepository) IsProductExist(productId int) (bool, error) {
// 	cart, err := cRepository.GetCart()
// 	if err != nil {
// 		return false, err
// 	}

// 	productList := []entity.CartProducts{}

// 	if err := cRepository.db.Where("cart_id = ?", cart.ID).Find(&productList).Error; err != nil {
// 		return false, err
// 	}

// 	for _, entry := range productList {
// 		if entry.ProductID == productId {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }

// func (cRepository CartRepository) RemoveProduct(productId int) error {
// 	cart, err := cRepository.GetCart()
// 	if err != nil {
// 		return err
// 	}
// 	productList := cart.Products
// 	var index int
// 	for i, product := range productList {
// 		if product.ID == productId {
// 			index = i
// 			break
// 		}
// 	}
// 	cRepository.db.Preload("Products").First(&cart)
// 	cart.TotalPrice -= cart.Products[index].Price
// 	cart.Products = append(cart.Products[:index], cart.Products[index+1:]...)
// 	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (cRepository CartRepository) GetProductsList() ([]entity.ProductEntry, error) {
// 	// cart, _ := cRepository.GetCart()
// 	productsList := []entity.ProductEntry{}
// 	if err := cRepository.db.Find(&productsList).Error; err != nil {
// 		return nil, err
// 	}
// 	return productsList, nil
// }
