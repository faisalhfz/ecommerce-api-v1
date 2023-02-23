package entity

type Cart struct {
	ID         int     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoincrement"`
	OrdersList []Order `json:"orders_list" gorm:"many2many:cart_orders;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice int     `json:"total_price"`
	IsCheckout bool    `json:"is_checkout"`
}
