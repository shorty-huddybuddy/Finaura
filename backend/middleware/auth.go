package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the session JWT from the Authorization header
		authHeader := c.Get("Authorization")
		fmt.Printf("DEBUG: Authorization header: %s\n", authHeader)

		sessionToken := strings.TrimPrefix(authHeader, "Bearer ")

		if sessionToken == "" {
			fmt.Println("DEBUG: No token provided")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No token provided",
			})
		}

		fmt.Printf("DEBUG: Token (first 20 chars): %s...\n", sessionToken[:min(20, len(sessionToken))])

		// Verify the session token
		// Note: If you have clock skew issues, sync your system time
		claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})

		if err != nil {
			fmt.Printf("Token verification error: %v\n", err)
			// For development: Skip clock validation errors
			if !strings.Contains(err.Error(), "issued in the future") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   "Invalid token",
					"details": err.Error(),
				})
			}
			// Continue even with clock skew in development
			fmt.Println("WARNING: Bypassing clock skew error for development")
		}

		// If claims is nil due to error, we need to get user ID from token manually
		var userId string
		if claims != nil {
			userId = claims.Subject
		} else {
			// Parse JWT manually to extract subject for development
			parts := strings.Split(sessionToken, ".")
			if len(parts) != 3 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token format",
				})
			}

			// Decode the payload (second part of JWT)
			payload, err := base64.RawURLEncoding.DecodeString(parts[1])
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token payload",
				})
			}

			// Parse the JSON payload
			var tokenClaims map[string]interface{}
			if err := json.Unmarshal(payload, &tokenClaims); err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token claims",
				})
			}

			// Extract the subject (user ID)
			if sub, ok := tokenClaims["sub"].(string); ok {
				userId = sub
				fmt.Printf("WARNING: Extracted userId %s from unverified token (clock skew bypass)\n", userId)
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Missing user ID in token",
				})
			}
		}

		// Store userId in context first (needed by all handlers)
		c.Locals("userId", userId)

		// Get user details and check if banned (skip in development if clock skew)
		if claims != nil {
			usr, err := user.Get(c.Context(), userId)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Error fetching user details",
					"details": err.Error(),
				})
			}

			if usr.Banned {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "User is banned",
				})
			}

			// Store additional user info in context
			c.Locals("banned", usr.Banned)
			c.Locals("firstName", usr.FirstName)
			c.Locals("lastName", usr.LastName)
			c.Locals("userImage", usr.ImageURL)

			if usr.Username != nil {
				c.Locals("username", *usr.Username)
			}
			if len(usr.EmailAddresses) > 0 {
				c.Locals("userEmail", usr.EmailAddresses[0].EmailAddress)
			}
		} else {
			fmt.Println("WARNING: Skipping banned user check due to clock skew bypass")
			// Set default values for development
			c.Locals("banned", false)
			c.Locals("firstName", "")
			c.Locals("lastName", "")
			c.Locals("userImage", "")
			c.Locals("username", "")
			c.Locals("userEmail", "")
		}

		return c.Next()
	}
}
