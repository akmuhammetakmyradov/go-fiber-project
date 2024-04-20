package middlewares

import (
	"errors"
	"fmt"
	"time"

	"github.com/akmuhammetakmyradov/test/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type LocalUser struct {
	ID   float64
	Type string
}

func TokenClaims(token, secret string) (*jwt.MapClaims, error) {

	decoded, err := jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !(ok && decoded.Valid) {
		return nil, errors.New("invalid token")
	}

	if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return &claims, nil
}

func MiddTokenChkUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Authorization header missing",
		})
	}

	config, _ := config.LoadConfiguration()

	claims, err := TokenClaims(token, config.JWT.AccessSecret)

	if err != nil || (*claims)["id"] == "" {
		fmt.Printf("Err in TokenClaims: %v", err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	c.Locals("user", LocalUser{
		ID:   (*claims)["id"].(float64),
		Type: (*claims)["type"].(string),
	})

	return c.Next()
}

func MiddChkAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(LocalUser)
	if user.Type != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "you don't have admin access",
		})
	}
	return c.Next()
}

func TokenEncode(claims *jwt.MapClaims, secret string, expiryAfter int64) (string, error) {

	// default 1 hour
	if expiryAfter == 0 {
		expiryAfter = 60 * 60
	}

	// or you can use time.Now().Add(time.Second * time.Duration(expiryAfter)).UTC().Unix()
	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// our signed jwt token string
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("error creating a token")
	}

	return signedToken, nil
}
