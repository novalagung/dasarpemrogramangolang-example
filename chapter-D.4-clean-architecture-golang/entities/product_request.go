package entities

type CreateProductRequest struct {
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}

type UpdateProductRequest struct {
	ID     int     `json:"id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}

type DeleteProductRequest struct {
	ID int `json:"id" validate:"required"`
}
