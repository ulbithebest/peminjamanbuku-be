package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/peminjamanbuku-be/model"
	"github.com/ulbithebest/peminjamanbuku-be/query"
	"gorm.io/gorm"
	"net/http"
)

func GetAllBuku(c *fiber.Ctx) error {
	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk mendapatkan semua buku
	books, err := query.GetAllBuku(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data buku, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada buku yang ditemukan, mengembalikan pesan kesalahan
	if len(books) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data buku tidak ditemukan", "data": nil})
	}

	// Jika tidak ada kesalahan, mengembalikan data buku sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    books,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func GetBukuByID(c *fiber.Ctx) error {
	// Mendapatkan parameter ID buku dari URL
	id := c.Params("id_buku")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk mendapatkan buku berdasarkan ID
	buku, err := query.GetBukuByID(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada kesalahan, mengembalikan data buku sebagai respons JSON
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": buku})
}

func InsertBuku(c *fiber.Ctx) error {
	// Mendeklarasikan variabel untuk menyimpan data buku dari body request
	var buku model.Buku

	// Mem-parsing body request ke dalam variabel buku
	if err := c.BodyParser(&buku); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk menyisipkan data buku ke dalam database
	if err := query.InsertBuku(db, buku); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan buku"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Buku berhasil disimpan", "data": buku})
}

func UpdateBuku(c *fiber.Ctx) error {
	// Mendapatkan parameter ID
	id := c.Params("id_buku")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	// Mendeklarasikan variabel untuk menyimpan data buku yang diperbarui dari body request
	var updatedBuku model.Buku

	// Mem-parsing body request ke dalam variabel updatedBuku
	if err := c.BodyParser(&updatedBuku); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk memperbarui data buku di dalam database
	if err := query.UpdateBuku(db, id, updatedBuku); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui buku"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Buku berhasil diperbarui"})
}

func DeleteBuku(c *fiber.Ctx) error {
	// Mendapatkan parameter ID buku dari URL
	id := c.Params("id_buku")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID buku tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi query untuk menghapus data buku dari database berdasarkan ID
	if err := query.DeleteBuku(db, id); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus buku"})
	}

	// Mengembalikan respons sukses
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Buku berhasil dihapus", "deleted_id": id})
}
