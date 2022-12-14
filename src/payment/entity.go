package payment

import (
	"go/src/product"
	"time"
)

type Payment struct {
	ID        int `gorm:"primaryKey"`
	ProductId int
	Product   product.Product `gorm:"foreignKey:ProductId"`
	PricePaid float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
