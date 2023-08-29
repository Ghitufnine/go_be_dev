package auth

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

func AccessSecretVersion(name string) string {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	var secretKey string
	secretKey = string(result.Payload.Data)

	return secretKey
}
func GenerateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func PostDoLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	postDoLogin := PostDoLoginQuery(username, password)
	if len(postDoLogin) > 0 {
		// Generate JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username // You can add more claims like user ID, roles, etc.
		// Add expiration time if needed: claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Sign and get the complete encoded token as a string
		name := "projects/77840544412/secrets/fleetify_api_jwt_key/versions/latest"
		secretKey := AccessSecretVersion(name)
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"data":  postDoLogin,
			"token": tokenString,
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid",
			"status":  0,
		})
	}
}

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello World!",
	})
}

func GetUserInfoFromToken(tokenString string) (jwt.MapClaims, error) {
	name := "projects/77840544412/secrets/fleetify_api_jwt_key/versions/latest"
	secretKey := AccessSecretVersion(name)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	return claims, nil
}
