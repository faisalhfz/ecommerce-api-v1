package request

import "ecommerce-api/src/entity"

type CreateCartRequest struct {
	OrdersList []*entity.Order `json:"orders_list"`
	TotalPrice int             `json:"total_price"`
	IsCheckout bool            `json:"is_checkout"`
}
