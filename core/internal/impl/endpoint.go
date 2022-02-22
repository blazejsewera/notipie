package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/jazzsewera/notipie/core/internal/impl/net"
	"github.com/jazzsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
)

type Endpoint struct {
	r         *gin.Engine
	grid      grid.Grid
	appProxy  grid.AppProxy
	userProxy grid.UserProxy
	hub       *net.Hub
	l         *zap.Logger
}

func NewEndpoint(grid grid.Grid, appProxy grid.AppProxy, userProxy grid.UserProxy) *Endpoint {
	return &Endpoint{
		r:         gin.Default(),
		grid:      grid,
		appProxy:  appProxy,
		userProxy: userProxy,
		l:         log.For("impl").Named("endpoint"),
	}
}

func (e *Endpoint) Setup() {
	e.r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	e.r.OPTIONS("/push", net.PreflightHandler)
	e.r.POST("/push", net.PushNotificationHandlerFor(e.appProxy))

	e.r.GET("/ws", net.WSHandlerFor(e.hub))
}

func (e *Endpoint) Run() {
	err := e.r.Run()
	if err != nil {
		e.l.Fatal("could not run endpoint", zap.Error(err))
		return
	}
	e.l.Info("endpoint running")
}
