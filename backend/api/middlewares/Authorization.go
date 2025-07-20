package middlewares

import (
	"strings"

	"slices"

	utils "github.com/AlvinSetyaPranata/ZENITH/backend/utils/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthenticationCheck(c *fiber.Ctx) error {
	tokenHeader := c.Get("Authorization")

	if tokenHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messege": "Unauthorized, You have no access rights!",
		})
	}

	splitedTokenHeader := strings.Split(tokenHeader, " ")

	if len(splitedTokenHeader) != 2 || splitedTokenHeader[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messege": "Unauthorized, You have no access rights!",
		})
	}

	token := splitedTokenHeader[1]

	// fmt.Println(token)

	userID, role, isValid := utils.ParseToken(token, true)

	if isValid != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messege": "Unauthorized, You have no access rights!",
		})
	}

	c.Locals("userID", userID)
	c.Locals("role", role)

	return c.Next()
}

func RoleGuard(allowedRoles ...string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)

		if slices.Contains(allowedRoles, role) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"messege": "You, forbids to access this resources",
		})
	}

}
