package product

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Price     float64   `gorm:"type:float"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
