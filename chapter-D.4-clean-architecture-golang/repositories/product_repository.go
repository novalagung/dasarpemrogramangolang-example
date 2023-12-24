package repositories

import (
	"clean-architecture-golang-example/entities"
	"errors"
	"log"

	"gorm.io/gorm"
)

// ProductRepositores: repository untuk model product
type ProductRepositories struct {
	database *gorm.DB
}

// NewProductRepositories: Injeksi repository product model
func NewProductRepositories(conn *gorm.DB, isMigrate bool) entities.ProductRepository {
	if isMigrate {
		err := conn.AutoMigrate(entities.Products{})
		if err != nil {
			log.Fatal("Migration Error:", err)
		}
	}
	return &ProductRepositories{conn}
}

// Create: digunakan untuk membuat insert data ke model product.
func (p *ProductRepositories) Create(product *entities.Products) error {
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(entities.Products{}).Create(product)
	err = query.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// Count: digunakan untuk menghitung jumlah data product yang tersimpan.
func (p *ProductRepositories) Count() (int64, error) {
	var count int64
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(entities.Products{}).Select("*").Count(&count)
	err = query.Error
	if err != nil {
		return count, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return count, err
	}

	return count, err
}

// DelteByID: digunakan untuk menghapus data product dengan id yang dipilih.
func (p *ProductRepositories) DeleteByID(id int) (entities.Products, error) {
	var product entities.Products
	var err error

	queryFind := p.database.Model(entities.Products{}).Where("id = ?", id).Find(&product)
	err = queryFind.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("ID is not found")
		}
		return product, err
	}

	queryDelete := queryFind.Delete(&product)
	err = queryDelete.Error
	if err != nil {
		return product, err
	}

	return product, err
}

// GetByID: digunakan untuk menampilkan data product yang sesuai dengan id yang dipilih.
func (p *ProductRepositories) GetByID(id int) (entities.Products, error) {
	var result entities.Products
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(&entities.Products{}).Where("id = ?", id).Where("deleted_at IS NULL").First(&result)
	err = query.Error
	if err != nil {
		return result, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return result, err
	}

	return result, err
}

// Gets: digunakan untuk menampilkan semua data product.
func (p *ProductRepositories) Gets() ([]entities.Products, error) {
	var results []entities.Products
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(&entities.Products{}).Select("*").Where("deleted_at IS NULL").Find(&results)
	err = query.Error
	if err != nil {
		return results, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return results, err
	}

	return results, err
}

// Update: digunakan untuk update data product.
func (p *ProductRepositories) Update(product *entities.Products) error {
	var err error
	var tx *gorm.DB = p.database.Begin()

	queryFind := tx.Model(entities.Products{}).Where("id = ?", product.ID).Updates(&product)
	err = queryFind.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("ID is not found")
		}
		return err
	}

	queryFind = tx.Commit()
	err = queryFind.Error
	if err != nil {
		return err
	}

	return err
}
