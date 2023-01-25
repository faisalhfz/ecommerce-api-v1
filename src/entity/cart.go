package entity

type Cart struct {
	ID           int            `json:"id" gorm:"column:id;type:bigint;primaryKey;autoincrement"`
	ProductsList []ProductEntry `json:"products_list" gorm:"many2many:cart_product_entries;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice   int            `json:"total_price"`
	IsCheckout   bool           `json:"is_checkout"`
}
