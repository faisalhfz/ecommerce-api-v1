package response

type GetCartResponse struct {
	ID         int                 `json:"cart_id"`
	OrdersList []*GetOrderResponse `json:"orders_list"`
	TotalPrice int                 `json:"total_price"`
}

type CartIDResponse struct {
	CartID int `json:"cart_id"`
}
