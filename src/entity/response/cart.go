package response

type GetCartResponse struct {
	ID           int                       `json:"cart_id"`
	ProductsList []GetProductEntryResponse `json:"products_list"`
	TotalPrice   int                       `json:"total_price"`
}

type GetProductEntryResponse struct {
	Product  GetProductResponse `json:"product"`
	Quantity int                `json:"quantity"`
}

type CartIDResponse struct {
	CartID int `json:"cart_id"`
}
