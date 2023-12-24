package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func NewDatabaseConfig(config DatabaseConfig) (db *gorm.DB, err error) {
	stringConnection := "host=" + config.Host + " user=" + config.Username + " password=" + config.Password + " dbname=" + config.DBname + " port=" + config.Port + " TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	// SetConnMaxIdleTime: berfungsi untuk menetapkan jumlah maksimum waktu koneksi secara idle(tidak berjalan)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
	// SetConnMaxLifetime: berfungsi menentukan maksimum waktu dapat digunakan kembali
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	// SetMaxIdleConns: berfungsi untuk menentukan jumlah connection tidak dijalankan(idle)
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns: berfungsi untuk menetapkan jumlah open koneksi
	sqlDB.SetMaxOpenConns(5)

	return
}
