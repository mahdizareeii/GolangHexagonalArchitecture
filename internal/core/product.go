package core

type Product struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	CreatedBy   string // User Id
}
