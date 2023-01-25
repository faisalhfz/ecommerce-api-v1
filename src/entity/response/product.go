package response

type GetProductResponse struct {
	ID       int    `json:"product_id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type GetProductDetailResponse struct {
	ID          int    `json:"product_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type ProductIDResponse struct {
	ProductID int `json:"product_id"`
}
