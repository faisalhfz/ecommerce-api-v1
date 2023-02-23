package request

type CreateProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ProductFilterRequest struct {
	Category string `json:"category" default:""`
	MinPrice int    `json:"minprice" default:"0"`
	MaxPrice int    `json:"maxprice" default:"0"`
}
