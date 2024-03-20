package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ulbithebest/peminjamanbuku-be/config"
	"github.com/ulbithebest/peminjamanbuku-be/routes"
)

func main() {
	// Membuat aplikasi Fiber
	app := fiber.New()

	// Koneksi ke database
	db := config.CreateDBConnection()

	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	// Menyimpan koneksi database dalam context Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Menetapkan rute untuk handler buku
	routes.SetupBookRoutes(app)

	// Menjalankan server Fiber
	app.Listen(":3000")
}
