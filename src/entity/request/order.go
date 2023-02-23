package request

import "ecommerce-api/src/entity"

type CreateOrderRequest struct {
	Product  *entity.Product `json:"product"`
	Quantity int             `json:"quantity"`
}
