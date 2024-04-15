package entity

type Product struct {
	ProductId               int     `gorm:"column:product_id;primaryKey" json:"productId"`
	ProductName             string  `gorm:"column:product_name;uniqueIndex" json:"productName"`
	ProductValue            float64 `gorm:"column:product_value" json:"productValue"`
	ProductDiscountPercent  float64 `gorm:"column:product_discount_percent" json:"productDiscountPercent"`
	ProductDescription      string  `gorm:"column:product_description" json:"productDescription"`
	ProductKgWeight         float64 `gorm:"column:product_kg_weitght" json:"productKgWeight"`
	ProductCentimeterWidth  float64 `gorm:"column:product_cm_width" json:"productCentimeterWidth"`
	ProductCentimeterHeight float64 `gorm:"column:product_centimeter_height" json:"productCentimeterHeight"`
	ProductCentimeterLength float64 `gorm:"column:product_cm_length" json:"productCentimeterLength"`

	ProductSellerId int  `gorm:"column:product_seller_id" json:"sellerId"`
	User            User `gorm:"foreignKey:ProductSellerId" json:"omitempty"`
}

func (Product) TableName() string {
	return "product"
}

func (p Product) Exist() bool {
	return p.ProductId != 0
}
