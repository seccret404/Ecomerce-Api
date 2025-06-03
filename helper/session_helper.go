package helper

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetSessionID(c *fiber.Ctx) string {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		// Jika tidak ada cookie, buat session ID baru
		sessionID = generateSessionID() 
		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			HTTPOnly: true,
		})
	}
	return sessionID
}

func generateSessionID() string {
	//generate session
	return "session_" + fmt.Sprintf("%d", time.Now().UnixNano())
}