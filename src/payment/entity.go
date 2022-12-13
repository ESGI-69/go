package payment

import (
	"time"
)

type Payment struct {
	ID        int `gorm:"primaryKey"`
	ProductId int
	Product   Product `gorm:"foreignKey:ProductId"`
	PricePaid float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID int `gorm:"primaryKey"`
}
