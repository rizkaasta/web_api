package controller

import (
	"net/http"
	"github.com/rizkaasta/web_api/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	ID				int		`json:"id"`
	Nama 			string	`json:"nama"`
	Prodi 			string	`json:"prodi"`
	Fakultas 		string	`json:"fakultas"`
	NIM 			int		`json:"nim"`
	TahunAngkatan 	int		`json:"tahun"`
}

//Create Data
func CreateDataMhs(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//validasi inputan
	var dataInput MahasiswaInput
	// if err := c.ShouldBindJSON(&dataInput.Nama);
	// err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error" : "Nama harus lebih dari 5 karakter",
	// 	})
	// 	return
	// } else if err := c.ShouldBindJSON(&dataInput.NIM);
	// err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error" : "NIM harus angka dan lebih dari 5 angka",
	// 	})
	// 	return
	// } else if err := c.ShouldBindJSON(&dataInput.TahunAngkatan);
	// err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error" : "Tahun angkatan harus berupa angka",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
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