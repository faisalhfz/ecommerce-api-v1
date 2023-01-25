package request

import "ecommerce-api/src/entity"

type CreateProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type CreateProductEntryRequest struct {
	Product  entity.Product `json:"product"`
	Quantity int            `json:"quantity"`
}

type ProductFilterRequest struct {
	Category string `json:"category" default:""`
	MinPrice int    `json:"minprice" default:"0"`
	MaxPrice int    `json:"maxprice" default:"0"`
}
