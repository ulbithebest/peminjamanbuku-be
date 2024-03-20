package query

import (
	"github.com/ulbithebest/peminjamanbuku-be/model"
	"gorm.io/gorm"
)

func GetAllBuku(db *gorm.DB) ([]model.Buku, error) { // Mengambil semua data buku dari database
	var buku []model.Buku
	if err := db.Find(&buku).Error; err != nil {
		return nil, err
	}
	return buku, nil
}

func GetBukuByID(db *gorm.DB, id string) (model.Buku, error) { // Mengambil data buku berdasarkan ID dari database
	var buku model.Buku
	if err := db.First(&buku, id).Error; err != nil {
		return buku, err
	}
	return buku, nil
}

func InsertBuku(db *gorm.DB, buku model.Buku) error { // Insert data buku ke dalam database
	if err := db.Create(&buku).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBuku(db *gorm.DB, id string, updatedBuku model.Buku) error { // Memperbarui data buku dalam database berdasarkan ID
	if err := db.Model(&model.Buku{}).Where("id_buku = ?", id).Updates(updatedBuku).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBuku(db *gorm.DB, id string) error { // Menghapus data buku dari database berdasarkan ID
	if err := db.Delete(&model.Buku{}, id).Error; err != nil {
		return err
	}
	return nil
}
