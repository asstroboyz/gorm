package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Membuka koneksi ke database PostgreSQL
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres dbname=go_restapi_bin port=5432 sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		// Jika terjadi kesalahan saat membuka koneksi, hentikan aplikasi dan tampilkan pesan kesalahan
		panic(err)
	}

	// Pastikan struktur model "Product" sesuai dengan tabel yang ada di database Anda.
	// Melakukan migrasi otomatis untuk memastikan struktur tabel yang sesuai dengan model tersedia
	database.AutoMigrate(&Product{})

	// Menetapkan instance DB untuk digunakan oleh package lain
	DB = database
}
