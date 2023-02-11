package entity

import "github.com/shopspring/decimal"

type Products struct {
	Model
	Name        string          `json:"name" gorm:"column:name"`
	Description string          `json:"description" gorm:"column:description"`
	Price       decimal.Decimal `json:"price" gorm:"column:price;type:decimal(10, 2)"`
	CategoryId  int             `json:"category_id" gorm:"column:category_id"`
	BrandId     int             `json:"brand_id" gorm:"column:brand_id"`
}
