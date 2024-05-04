package entity

type ShoppingCart struct {
	ShoppingCartId           int `gorm:"column:shopping_cart_id;primaryKey" json:"shoppingCartId,omitempty"`
	ShoppingCartItemQuantity int `gorm:"column:shopping_cart_item_quantity" json:"shoppingCartItemQuantity,omitempty"`

	ShoppingCartItemId int     `gorm:"column:shopping_cart_item_id" json:"shoppingCartItemId,omitempty"`
	Product            Product `gorm:"foreignKey:ShoppingCartItemId" json:"product,omitempty"`

	ShoppingCartUserId int  `gorm:"column:shopping_cart_user_id" json:"shoppintCartUserId,omitempty"`
	User               User `gorm:"foreignKey:ShoppingCartUserId" json:"user,omitempty"`
}
