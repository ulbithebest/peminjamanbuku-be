package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/peminjamanbuku-be/controller"
)

func SetupBookRoutes(app *fiber.App) {
	app.Get("/buku", controller.GetAllBuku)
	app.Get("/buku/getbyid/:id_buku", controller.GetBukuByID)
	app.Post("buku/insert", controller.InsertBuku)
	app.Put("/buku/update/:id_buku", controller.UpdateBuku)
	app.Delete("/buku/delete/:id_buku", controller.DeleteBuku)
}

func SetupAnggotaRoutes(app *fiber.App) {
	app.Get("/anggota", controller.GetAllAnggota)
	app.Get("/anggota/getbyid/:id_anggota", controller.GetAnggotaById)
	app.Post("anggota/insert", controller.InsertAnggota)
	app.Put("/anggota/update/:id_anggota", controller.UpdateAnggota)
	app.Delete("/anggota/delete/:id_anggota", controller.DeleteAnggota)
}
