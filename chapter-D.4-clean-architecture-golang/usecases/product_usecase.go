package usecases

import (
	"clean-architecture-golang-example/entities"
	"errors"

	"github.com/go-playground/validator/v10"
)

// ProductUsecase: sebagai orchestrator bisnis proses product.
type ProductUsecase struct {
	repository *entities.ProductRepository
	valid      *validator.Validate
}

// NewProductUsecase: injeksi dari repository ke usecase
func NewProductUsecase(repository *entities.ProductRepository, valid *validator.Validate) entities.ProductUsecase {
	return &ProductUsecase{repository, valid}
}

// Create: digunakan untuk insert product ke repository.
func (usecase *ProductUsecase) Create(product entities.CreateProductRequest) (entities.ProductResponseJSON, error) {
	var err error
	var result entities.ProductResponseJSON
	repo := *usecase.repository
	err = usecase.valid.Struct(product)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error validation:" + err.Error(),
		}
		return result, err
	}

	count, err := repo.Count()
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	var data = entities.Products{
		ID:     int(count) + 1,
		Name:   product.Name,
		Price:  product.Price,
		Weight: product.Weight,
	}

	if err = repo.Create(&data); err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   1,
		Success: true,
		Message: "Create product success",
	}

	return result, err
}

// DeleteByID: digunakan untuk hapus product dengan id ke repository.
func (usecase *ProductUsecase) DeleteByID(id int) (entities.ProductResponseJSON, error) {
	var result entities.ProductResponseJSON
	if id == 0 {
		return result, errors.New("ID must be not empty")
	}
	repo := *usecase.repository
	data, err := repo.DeleteByID(id)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   1,
		Success: true,
		Message: "Delete product success",
	}

	return result, nil
}

// GetOne: digunakan untuk mengambil data product dengan id yang sudah dipilih
func (usecase *ProductUsecase) GetOne(id int) (entities.ProductResponseJSON, error) {
	var result entities.ProductResponseJSON
	if id == 0 {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server: ID must be not empty",
		}
		return result, errors.New("ID must be not empty")
	}
	repo := *usecase.repository

	data, err := repo.GetByID(id)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server: " + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data: []entities.Products{
			data,
		},
		Count:   1,
		Success: true,
	}

	return result, nil
}

// Gets: digunakan untuk menampilkan semua data product
func (usecase *ProductUsecase) Gets() (entities.ProductResponseJSON, error) {
	repo := *usecase.repository
	data, err := repo.Gets()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	count, err := repo.Count()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	result := entities.ProductResponseJSON{
		Data:    data,
		Count:   count,
		Success: true,
	}

	return result, nil
}

// Update: digunakan untuk mengubah data dengan id yang sudah dipilih
func (usecase *ProductUsecase) Update(product entities.UpdateProductRequest) (entities.ProductResponseJSON, error) {
	err := usecase.valid.Struct(product)
	var result entities.ProductResponseJSON
	if err != nil {
		return result, err
	}
	repo := *usecase.repository
	var data = entities.Products{
		ID:     product.ID,
		Name:   product.Name,
		Price:  product.Price,
		Weight: product.Weight,
	}

	if err = repo.Update(&data); err != nil {
		return result, err
	}

	count, err := repo.Count()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   count,
		Success: true,
		Message: "Update product success",
	}

	return result, nil
}
