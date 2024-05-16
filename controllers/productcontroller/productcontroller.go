package productcontroller

import (
	"encoding/json"
	"go_restapi_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	//mengambil semua produk di database
	models.DB.Find(&products)

	// mengembalikan daftar produk sebagai json
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	//mencari produk berdasarkan id
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			// Jika produk tidak ditemukan, kembalikan respons status 404
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			// Jika terjadi kesalahan lain, kembalikan respons status 500
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan server"})
			return
		}
	}
	// Mengembalikan detail produk sebagai respons JSON
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product
	// Mendekode data JSON yang diterima menjadi struktur produk
	if err := c.ShouldBindJSON(&product); err != nil {

		// Jika data tidak valid, kembalikan respons status 400
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Menyimpan produk baru ke dalam database
	models.DB.Create(&product)

	// Mengembalikan detail produk yang baru dibuat sebagai respons JSON
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	// Mendekode data JSON yang diterima menjadi struktur produk
	if err := c.ShouldBindJSON(&product); err != nil {

		// Jika data tidak valid, kembalikan respons status 400
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Memperbarui produk dalam database berdasarkan ID
	result := models.DB.Model(&models.Product{}).Where("id = ?", id).Updates(&product)
	if result.Error != nil {

		// Jika terjadi kesalahan dalam proses pembaruan, kembalikan respons status 500
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan server"})
		return
	}

	// Jika produk dengan ID yang diberikan tidak ditemukan, kembalikan respons status 400
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat memperbarui data"})
		return
	}

	models.DB.Find(&product)
	// Mengembalikan pesan sukses sebagai respons JSON
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui", "product": product})

}

func Delete(c *gin.Context) {
	var product models.Product

	// input := map[string]string{"id": "0"}
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus produk"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
