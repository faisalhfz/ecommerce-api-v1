package entity

type Product struct {
	ID          int    `json:"id" gorm:"column:id;type:bigint;primaryKey;autoincrement"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ProductEntry struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}
