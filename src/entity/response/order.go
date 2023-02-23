package response

type GetOrderResponse struct {
	ID       int                 `json:"order_id"`
	Product  *GetProductResponse `json:"product"`
	Quantity int                 `json:"quantity"`
}

type NewOrderResponse struct {
	CartID  int `json:"cart_id"`
	OrderID int `json:"order_id"`
}
