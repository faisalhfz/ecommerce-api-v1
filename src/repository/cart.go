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
	CheckoutCart() (*entity.Cart, error)
	GetCompletedCarts() ([]entity.Cart, error)
	DeleteCart(cart *entity.Cart) error
	AddOrder(order *entity.Order) (int, error)
	RemoveOrder(order *entity.Order) error
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
	cRepository.db.Model(&cart).Association("OrdersList").Append(&cart.OrdersList)
	return cart.ID, nil
}

func (cRepository CartRepository) GetCart() (*entity.Cart, error) {
	var cart *entity.Cart
	if err := cRepository.db.Preload("OrdersList").Where("is_checkout = ?", false).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (cRepository CartRepository) CheckoutCart() (*entity.Cart, error) {
	cart, err := cRepository.GetCart()
	if err != nil {
		return nil, err
	}
	cRepository.db.Preload("OrdersList").First(&cart)
	cart.IsCheckout = true
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (cRepository CartRepository) GetCompletedCarts() ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := cRepository.db.Preload("OrdersList").Where("is_checkout = ?", true).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (cRepository CartRepository) DeleteCart(cart *entity.Cart) error {
	cRepository.db.Model(&cart).Association("OrdersList").Delete(&cart.OrdersList)
	if err := cRepository.db.Debug().Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (cRepository CartRepository) AddOrder(order *entity.Order) (int, error) {
	cart, _ := cRepository.GetCart()

	cart.OrdersList = append(cart.OrdersList, *order)
	cart.TotalPrice += order.Product.Price * order.Quantity

	cRepository.db.Model(&cart).Association("OrdersList").Replace(&cart.OrdersList)
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func (cRepository CartRepository) RemoveOrder(order *entity.Order) error {
	cart, _ := cRepository.GetCart()

	var index int
	for i, entry := range cart.OrdersList {
		if entry.ProductID == order.ProductID {
			index = i
		}
	}
	cart.OrdersList = append(cart.OrdersList[:index], cart.OrdersList[index+1:]...)
	cart.TotalPrice -= order.Product.Price * order.Quantity

	if len(cart.OrdersList) == 0 {
		if err := cRepository.DeleteCart(cart); err != nil {
			return err
		}
		return nil
	}

	cRepository.db.Model(&cart).Association("OrdersList").Replace(&cart.OrdersList)
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return err
	}
	return nil
}
