package main

import (
    "github.com/gofiber/fiber/v2"
    "example.com/greetings/database"
    //"example.com/greetings/entity"
    "example.com/greetings/middleware"
    "example.com/greetings/controller"
    //"gorm.io/gen"
    //"log"
)

func init() {
	database.ConnectDB()

    sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

    // Automatically migrate all models
    /*models := []interface{}{
        &entity.AttendanceAdjustment{},
        &entity.AttendanceCredit{},
    }

    err = database.Migrate(database.DBConn, models...)
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }*/
    defer sqlDb.Close()

}

func main() {
    //init()


    // Inisialisasi GORM Gen
    /*g := gen.NewGenerator(gen.Config{
        OutPath: "./entity", // Ganti dengan path model yang Anda inginkan
        Mode:    gen.WithDefaultQuery, // Menghasilkan default query
    })

    // Menghubungkan ke database
    g.UseDB(database.DBConn)

    // Auto-generate model untuk semua tabel
    g.GenerateAllTable()

    // Atau untuk tabel tertentu
    // g.GenerateModel("users") // Ganti dengan nama tabel yang diinginkan

    // Menyimpan hasil
    g.Execute()

    return*/



	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "ok"})
		return c.SendString("Hello World.")
	})

    app.Post("/login", controller.Login)

    // Protected route
    app.Get("/protected", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
        user := c.Locals("user")
        return c.JSON(fiber.Map{"message": "Welcome", "user": user})
    })

	app.Listen(":8080")
}
