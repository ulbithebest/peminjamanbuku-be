package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/peminjamanbuku-be/model"
	"github.com/ulbithebest/peminjamanbuku-be/query"
	"gorm.io/gorm"
	"net/http"
)

func GetAllAnggota(c *fiber.Ctx) error {
	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk mendapatkan semua anggota
	anggota, err := query.GetAllAnggota(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data anggota, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada anggota yang ditemukan, mengembalikan pesan kesalahan
	if len(anggota) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data anggota tidak ditemukan", "data": nil})
	}

	// Jika tidak ada kesalahan, mengembalikan data anggota sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    anggota,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func GetAnggotaById(c *fiber.Ctx) error {
	// Mendapatkan parameter ID anggota dari URL
	id := c.Params("id_anggota")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID anggota tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk mendapatkan anggota berdasarkan ID
	anggota, err := query.GetAnggotaById(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada kesalahan, mengembalikan data anggota sebagai respons JSON
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": anggota})
}

func InsertAnggota(c *fiber.Ctx) error {
	// Mendeklarasikan variabel untuk menyimpan data anggota dari body request
	var anggota model.Anggota

	// Mem-parsing body request ke dalam variabel anggota
	if err := c.BodyParser(&anggota); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk menyisipkan data anggota ke dalam database
	if err := query.InsertAnggota(db, anggota); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan anggota"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Data Anggota berhasil disimpan", "data": anggota})
}

func UpdateAnggota(c *fiber.Ctx) error {
	// Mendapatkan parameter ID
	id := c.Params("id_anggota")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID anggota tidak ditemukan"})
	}

	// Mendeklarasikan variabel untuk menyimpan data anggota yang diperbarui dari body request
	var updatedAnggota model.Anggota

	// Mem-parsing body request ke dalam variabel updatedAnggota
	if err := c.BodyParser(&updatedAnggota); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk memperbarui data anggota di dalam database
	if err := query.UpdateAnggota(db, id, updatedAnggota); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui anggota"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Anggota berhasil diperbarui"})
}

func DeleteAnggota(c *fiber.Ctx) error {
	// Mendapatkan parameter ID anggota dari URL
	id := c.Params("id_anggota")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID anggota tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk menghapus data anggota dari database berdasarkan ID
	if err := query.DeleteAnggota(db, id); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus anggota"})
	}

	// Mengembalikan respons sukses
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Anggota berhasil dihapus", "deleted_id": id})
}
