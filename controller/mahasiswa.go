package controller

import (
	"net/http"
	"github.com/rizkaasta/web_api/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type MahasiswaInput struct {
	ID				int		`json:"id" binding:"required"`
	Nama 			string	`json:"nama" binding:"required,min=6"`
	Prodi 			string	`json:"prodi" binding:"required"`
	Fakultas 		string	`json:"fakultas" binding:"required"`
	NIM 			int		`json:"nim" binding:"required,gt=9999"`
	TahunAngkatan 	int		`json:"tahun" binding:"required"`
}

//Create Data
func CreateDataMhs(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//validasi inputan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			case "min":
				errorMessage := fmt.Sprintf("Error %s, message: nama harus terdiri dari 6 karakter atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			case "gt":
				errorMessage := fmt.Sprintf("Error %s, message: nim harus terdiri dari 6 angka atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}

	// input data
	mhs := models.Mahasiswa{
		ID				: dataInput.ID,			
		Nama 			: dataInput.Nama,
		Prodi 			: dataInput.Prodi,
		Fakultas 		: dataInput.Fakultas,
		NIM 			: dataInput.NIM,
		TahunAngkatan	: dataInput.TahunAngkatan,
	}
	db1.Create(&mhs)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Input data berhasil",
		"Data" : mhs,
		"Time" : time.Now(),
	})
}

//Read Data
func ReadDataMhs(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)
	
	var mhs []models.Mahasiswa
	db1.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"Data" : mhs,
		"Time" : time.Now(),
	})

}

//Update Data
func UpdateDataMhs(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db1.Where("nim = ?", c.Param("nim")).First(&mhs).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "min":
				errorMessage := fmt.Sprintf("Error %s, message: nama harus terdiri dari 6 karakter atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			case "gt":
				errorMessage := fmt.Sprintf("Error %s, message: nim harus terdiri dari 6 angka atau lebih", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}

	//ubah data
	db1.Model(&mhs).Update(&dataInput)

	//menampilkan data
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Data berhasil diubah",
		"Data" : mhs,
		"Time" : time.Now(),
	})
}

//Delete Data
func DeleteDataMhs(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db1.Where("nim = ?", c.Query("nim")).First(&mhs).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Data mahasiswa tidak di temukan",
		})
		return
	}

	//hapus data
	db1.Delete(&mhs)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data" : true,
	})
}