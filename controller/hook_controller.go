package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	
	"autochess/service"
	"autochess/entity"
)

// Controller is user controlller
type HookController struct{}

// Index action: GET /users
func (pc HookController) Index(c *gin.Context) {
	// var s service.HookService
	// s.CallBack("hello")
	c.JSON(200, gin.H{
		"message": "hello world!!!",
	})
}

// Verify Facebook Webhook
// GET: /hook
func (pc HookController) Verify(c *gin.Context) {
	verifyToken := "9BB23EAE99AE729AA66FF7D834149"
	// Parse query params
	mode := c.Request.URL.Query().Get("hub.mode")
	token := c.Request.URL.Query().Get("hub.verify_token")
	challenge := c.Request.URL.Query().Get("hub.challenge")

	log.Println("mode", mode, len(mode))
	if (len(mode) > 0 && len(token) > 0) {
		// Checks the mode and token sent is correct
		if (mode == "subscribe" && token == verifyToken) {
			// Responds with the challenge token from the request
			log.Println("WEBHOOK_VERIFIED")
			c.String(200, challenge)
		} else {
			// Responds with '403 Forbidden' if verify tokens do not match
			c.String(403, "Forbidden")
		}
	}
}

// Recieve facebook event
// POST: /hook
func (pc HookController) Event(c *gin.Context) {
	var hook entity.Hook
	if	c.BindJSON(&hook) == nil {
		var s service.HookService
		s.CallBack(
			hook.Entries[0].Messages[0].Sender.Id,
			hook.Entries[0].Messages[0].Message.Text,
		)
	}
	c.String(200, "EVENT_RECEIVED")
}