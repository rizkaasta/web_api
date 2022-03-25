package controller

import (
	"net/http"
	"github.com/rizkaasta/web_api/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MataKuliahInput struct {
	ID				 int		`json:"id"`
	Kode_MataKuliah	 string 	`json:"kode"`
	Nama_MataKuliah  string		`json:"nama"`
	JumlahSKS		 int		`json:"jumlahSKS"`
	DosenPengampu	 string		`json:"dosen"`
}

//Create Data
func CreateDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//validasi inputan
	var dataInput MataKuliahInput
	// if err := c.ShouldBindJSON(&dataInput.Nama_MataKuliah);
	// err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error" : "Nama mata kuliah harus lebih dari 3 karakter",
	// 	})
	// 	return
	// } else 
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Data required",
		})
		return
	}

	// input data
	matkul := models.MataKuliah{
		ID				: dataInput.ID,				 
		Kode_MataKuliah	: dataInput.Kode_MataKuliah,
		Nama_MataKuliah : dataInput.Nama_MataKuliah,
		JumlahSKS		: dataInput.JumlahSKS,
		DosenPengampu	: dataInput.DosenPengampu,
	}
	db2.Create(&matkul)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Input data berhasil",
		"Data" : matkul,
		"Time" : time.Now(),
	})
}

//Read Data
func ReadDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)
	
	var matkul []models.MataKuliah
	db2.Find(&matkul)
	c.JSON(http.StatusOK, gin.H{
		"Data" : matkul,
		"Time" : time.Now(),
	})
}

//Update Data
func UpdateDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db2.Where("kode = ?", c.Param("kode")).First(&matkul).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Mata kuliah tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MataKuliahInput
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	//ubah data
	db2.Model(&matkul).Update(&dataInput)

	//menampilkan data
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Data berhasil diubah",
		"Data" : matkul,
		"Time" : time.Now(),
	})
}

//Delete Data
func DeleteDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db2.Where("kode = ?", c.Query("kode")).First(&matkul).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Mata kuliah tidak di temukan",
		})
		return
	}

	//hapus data
	db2.Delete(&matkul)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data" : true,
	})
}