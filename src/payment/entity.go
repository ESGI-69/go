package payment

import (
	"go/src/product"
	"time"
)

type Payment struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Product   *product.Product `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE;foreignKey:ProductID"`
	PricePaid float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
