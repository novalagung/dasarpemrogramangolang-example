package entities

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID     int     `gorm:"column:id"`
	Name   string  `gorm:"column:name"`
	Price  float64 `gorm:"column:price"`
	Weight float64 `gorm:"column:weight"`
}

// Membuat model interface untuk repository product
type ProductRepository interface {
	GetByID(id int) (Products, error)
	Gets() ([]Products, error)
	Create(product *Products) error
	Update(product *Products) error
	DeleteByID(id int) (Products, error)
	Count() (int64, error)
}

// Membuat model interface untuk usecase product
type ProductUsecase interface {
	GetOne(id int) (ProductResponseJSON, error)
	Gets() (ProductResponseJSON, error)
	Create(product CreateProductRequest) (ProductResponseJSON, error)
	Update(product UpdateProductRequest) (ProductResponseJSON, error)
	DeleteByID(id int) (ProductResponseJSON, error)
}
