package services

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RemoveQueryAuthService(c *fiber.Ctx) error {
	queries := string(c.Request().URI().QueryString())

	arrQueries := strings.Split(queries, "&")
	newQueriesTemp := arrQueries[:0]

	for _, q := range arrQueries {
		if !strings.Contains(q, "oauth_") {
			newQueriesTemp = append(newQueriesTemp, q)
		}
	}
	queries = strings.Join(newQueriesTemp, "&")

	newUrl := c.Path()
	if queries != "" {
		newUrl = newUrl + "?" + queries
	}

	newUrl = strings.ReplaceAll(newUrl, "/wp-json/wc/v3/", "")

	c.Locals("newUrl", newUrl)

	return c.Next()
}
