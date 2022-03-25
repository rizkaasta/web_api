package models

type Mahasiswa struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Nama          string `json:"nama"`
	Prodi         string `json:"prodi"`
	Fakultas      string `json:"fakultas"`
	NIM           int    `json:"nim"`
	TahunAngkatan int    `json:"tahun"`
}

type MataKuliah struct {
	ID              int    `json:"id" gorm:"primary_key"`
	Kode_MataKuliah string `json:"kode"`
	Nama_MataKuliah string `json:"nama"`
	JumlahSKS       int    `json:"jumlahSKS"`
	DosenPengampu   string `json:"dosen"`
}
