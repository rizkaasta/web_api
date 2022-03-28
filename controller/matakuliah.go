package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/rizkaasta/web_api/models"
)

type MataKuliahInput struct {
	ID              int    `json:"id" binding:"required"`
	Kode_MataKuliah string `json:"kode" binding:"required"`
	Nama_MataKuliah string `json:"nama" binding:"required,min=4"`
	JumlahSKS       int    `json:"jumlahSKS" binding:"required"`
	DosenPengampu   string `json:"dosen" binding:"required"`
}

//Create Data
func CreateDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//validasi inputan
	var dataInput MataKuliahInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			case "min":
				errorMessage := fmt.Sprintf("Error %s, message: nama harus terdiri dari 6 karakter atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// input data
	matkul := models.MataKuliah{
		ID:              dataInput.ID,
		Kode_MataKuliah: dataInput.Kode_MataKuliah,
		Nama_MataKuliah: dataInput.Nama_MataKuliah,
		JumlahSKS:       dataInput.JumlahSKS,
		DosenPengampu:   dataInput.DosenPengampu,
	}
	db2.Create(&matkul)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Input data berhasil",
		"Data":    matkul,
		"Time":    time.Now(),
	})
}

//Read Data
func ReadDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	var matkul []models.MataKuliah
	db2.Find(&matkul)
	c.JSON(http.StatusOK, gin.H{
		"Data": matkul,
		"Time": time.Now(),
	})
}

//Update Data
func UpdateDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db2.Where("kode = ?", c.Param("kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Mata kuliah tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MataKuliahInput
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "min":
				errorMessage := fmt.Sprintf("Error %s, message: nama harus terdiri dari 6 karakter atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	//ubah data
	db2.Model(&matkul).Update(&dataInput)

	//menampilkan data
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data berhasil diubah",
		"Data":    matkul,
		"Time":    time.Now(),
	})
}

//Delete Data
func DeleteDataMatkul(c *gin.Context) {
	db2 := c.MustGet("db2").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db2.Where("kode = ?", c.Query("kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Mata kuliah tidak di temukan",
		})
		return
	}

	//hapus data
	db2.Delete(&matkul)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": true,
	})
}
