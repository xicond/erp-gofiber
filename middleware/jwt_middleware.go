package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    //"golang.org/x/crypto/bcrypt"
    "example.com/greetings/entity"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "log"
    "time"
)

var mockUser = entity.User{
    ID:       1,
    Email: "email@example.com",
    Password: "$2a$10$Xj1uJ5I3b5gTPxQ4z0N9EOQz5iD6dU5tZp0sVg0B5cG7YqF1Q5T6y", // bcrypt hashed password for "password"
}

// Load RSA keys
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
    privateKeyData, err := ioutil.ReadFile("private.key")
    if err != nil {
        log.Fatal("Error loading private key:", err)
    }

    block, _ := pem.Decode(privateKeyData)
    if block == nil || block.Type != "PRIVATE KEY" {
        log.Fatal("Failed to parse private key")
    }

    privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Fatal("Error parsing private key:", err)
    }

    publicKeyData, err := ioutil.ReadFile("public.key")
    if err != nil {
        log.Fatal("Error loading public key:", err)
    }

    block, _ = pem.Decode(publicKeyData)
    if block == nil || block.Type != "PUBLIC KEY" {
        log.Fatal("Failed to parse public key")
    }

    parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        log.Fatal("Error parsing public key:", err)
    }

    // Type assertion to *rsa.PublicKey
    var ok bool
    publicKey, ok = parsedKey.(*rsa.PublicKey)
    if !ok {
        log.Fatal("Public key is not of type *rsa.PublicKey")
    }
}

func GenerateJWT(userID uint) (string, error) {
    claims := &jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 1).Unix(), // Token valid for 1 hour
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(privateKey)
}

// ValidateJWT validates the token and returns the user ID
func ValidateJWT(tokenString string) (uint, error) {
    claims := &jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
        return publicKey, nil
    })

    if err != nil || !token.Valid {
        return 0, err
    }

    // Dereference claims to access the actual map
    userID := uint((*claims)["user_id"].(float64)) // Convert float64 to uint
    return userID, nil
}

// JWT middleware
func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing authorization header"})
        }

        // Remove "Bearer " prefix
        token = token[len("Bearer "):]

        // Validate the token
        username, err := ValidateJWT(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        // Store username in context
        c.Locals("user", username)
        return c.Next()
    }
}
