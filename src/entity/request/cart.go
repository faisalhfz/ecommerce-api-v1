package request

import "ecommerce-api/src/entity"

type CreateCartRequest struct {
	ProductsList []entity.ProductEntry `json:"products_list"`
	TotalPrice   int                   `json:"total_price"`
	IsCheckout   bool                  `json:"is_checkout"`
}
