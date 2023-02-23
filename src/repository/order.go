package repository

import (
	"ecommerce-api/src/entity"
	"ecommerce-api/src/entity/request"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(orderRequest request.CreateOrderRequest) (int, error)
	GetOrder(productId int) (*entity.Order, error)
	DeleteOrder(order *entity.Order) error
	// UpdateOrder(order *entity.Order, newOrder request.CreateProductEntryRequest) error
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (oRepository OrderRepository) CreateOrder(orderRequest request.CreateOrderRequest) (int, error) {
	var order entity.Order
	copier.Copy(&order, &orderRequest)
	oRepository.db.Model(&order).Association("Product").Append(&order.Product)
	if err := oRepository.db.Debug().Create(&order).Error; err != nil {
		return 0, err
	}
	return order.ID, nil
}

func (oRepository OrderRepository) GetOrder(id int) (*entity.Order, error) {
	var order *entity.Order
	if err := oRepository.db.Preload("Product").First(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (oRepository OrderRepository) DeleteOrder(order *entity.Order) error {
	if err := oRepository.db.Debug().Delete(&order).Error; err != nil {
		return err
	}
	return nil
}
