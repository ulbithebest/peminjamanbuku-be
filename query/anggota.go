package query

import (
	"github.com/ulbithebest/peminjamanbuku-be/model"
	"gorm.io/gorm"
)

func GetAllAnggota(db *gorm.DB) ([]model.Anggota, error) { // Mengambil semua data anggota dari database
	var anggota []model.Anggota
	if err := db.Find(&anggota).Error; err != nil {
		return nil, err
	}
	return anggota, nil
}

func GetAnggotaById(db *gorm.DB, id string) (model.Anggota, error) { // Mengambil data anggota berdasarkan ID dari database
	var anggota model.Anggota
	if err := db.First(&anggota, id).Error; err != nil {
		return anggota, err
	}
	return anggota, nil
}

func InsertAnggota(db *gorm.DB, anggota model.Anggota) error { // Insert data anggota ke dalam database
	if err := db.Create(&anggota).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAnggota(db *gorm.DB, id string, updatedAnggota model.Anggota) error { // Memperbarui data anggota dalam database berdasarkan ID
	if err := db.Model(&model.Anggota{}).Where("id_anggota = ?", id).Updates(updatedAnggota).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnggota(db *gorm.DB, id string) error { // Menghapus data anggota dari database berdasarkan ID
	if err := db.Delete(&model.Anggota{}, id).Error; err != nil {
		return err
	}
	return nil
}
