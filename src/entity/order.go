package entity

type Order struct {
	ID        int     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoincrement"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}
