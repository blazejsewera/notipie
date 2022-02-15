package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/jazzsewera/notipie/core/internal/impl/net"
)

type Endpoint struct {
	r *gin.Engine
}

func (e *Endpoint) SetupFor(hub *net.Hub) {
	e.r = gin.Default()

	e.r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	e.r.OPTIONS("/push", net.PreflightHandler)
	e.r.POST("/push", net.PushNotificationHandlerFor(hub))

	e.r.GET("/ws", net.WSHandlerFor(hub))
}

func (e *Endpoint) Run() {
	e.r.Run()
}
