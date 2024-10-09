package controller

import (
    "github.com/gofiber/fiber/v2"
    "example.com/greetings/entity"
    "example.com/greetings/middleware"
    "example.com/greetings/database"
    "golang.org/x/crypto/bcrypt"
)

/// Login handler
func Login(c *fiber.Ctx) error {
    var input entity.User
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Retrieve user from database
    var user entity.User
    if err := database.DBConn.Where("username = ?", input.Email).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Generate JWT token
    accessToken, err := middleware.GenerateJWT(uint(user.ID)) // Convert int64 to uint
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
    }

    return c.JSON(fiber.Map{"access_token": accessToken})
}
