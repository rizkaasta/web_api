package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetUpMahasiswa() *gorm.DB {
	db1, err := gorm.Open("mysql", "root:@(localhost)/dbmahasiswa?charset=utf8&parseTime=True&loc=Local")	
	if err != nil {
		panic("Error koneksi kedalam database")
	}

	db1.AutoMigrate(&Mahasiswa{})
	
	return db1
}

func SetUpMataKuliah() *gorm.DB {
	db2, err := gorm.Open("mysql", "root:@(localhost)/dbmatakuliah?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error koneksi kedalam database")
	}
	
	db2.AutoMigrate(&MataKuliah{})
	return db2
}