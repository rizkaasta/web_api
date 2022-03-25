package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rizkaasta/web_api/controller"
	"github.com/rizkaasta/web_api/models"
)

func main() {
	r := gin.Default()
	db1 := models.SetUpMahasiswa()
	db2 := models.SetUpMataKuliah()
	r.Use(func (c *gin.Context)  {
		c.Set("db1", db1)
		c.Set("db2", db2)
		c.Next()
	})

	v1 := r.Group("api/v1/")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message" : "Web API Universitas A",
			})
		})

		//Membuat data baru mahasiswa
		v1.POST("/mahasiswa", controller.CreateDataMhs)
		//Menampilkan semua data mahasiswa
		v1.GET("/mahasiswa", controller.ReadDataMhs)
		//Mengubah data mahasiswa
		v1.PUT("/mahasiswa/:nim", controller.UpdateDataMhs)
		//Menghapus data mahasiswa
		v1.DELETE("/mahasiswa", controller.DeleteDataMhs)

		//Membuat data baru mata kuliah
		v1.POST("/matakuliah", controller.CreateDataMatkul)
		//Menampilkan semua data mata kuliah
		v1.GET("/matakuliah", controller.ReadDataMatkul)
		//Mengubah data mata kuliah
		v1.PUT("/matakuliah/:kode", controller.UpdateDataMatkul)
		//Menghapus data mahasiswa
		v1.DELETE("/matakuliah", controller.DeleteDataMatkul)
	}

	r.Run()
}