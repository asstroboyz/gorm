package main

import (
	"go_restapi_gin/controllers/productcontroller"
	"go_restapi_gin/models"

	"github.com/gin-gonic/gin"
)

// main adalah titik masuk utama aplikasi.
func main() {
	// Membuat instance router Gin dengan konfigurasi default.
	r := gin.Default()

	// Menghubungkan ke database saat aplikasi dimulai.
	models.ConnectDatabase()

	// Mendefinisikan rute-rute untuk API produk.
	r.GET("/api/products", productcontroller.Index)         // Menampilkan daftar semua produk.
	r.GET("/api/products/:id", productcontroller.Show)      // Menampilkan detail produk berdasarkan ID.
	r.POST("/api/products", productcontroller.Create)       // Membuat produk baru.
	r.PUT("/api/products/:id", productcontroller.Update)    // Memperbarui produk berdasarkan ID.
	r.DELETE("/api/products/:id", productcontroller.Delete) // Menghapus produk berdasarkan ID.

	// Menjalankan server web pada port yang ditentukan dan meng-handle
	// permintaan HTTP yang masuk menggunakan router yang telah dikonfigurasi.
	r.Run()
}
