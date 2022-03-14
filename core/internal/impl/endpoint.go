package impl

import (
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/net"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Endpoint struct {
	r    *gin.Engine
	grid grid.Grid
	hub  *ws.HubImpl
	l    *zap.Logger
}

func NewEndpoint(grid grid.Grid) *Endpoint {
	return &Endpoint{
		r:    gin.New(),
		grid: grid,
		l:    log.For("impl").Named("endpoint"),
	}
}

func (e *Endpoint) Setup() {
	e.r.Use(gin.Recovery())
	e.r.GET("/", net.PingHandler)

	e.r.OPTIONS("/push", net.PreflightHandler)
	e.r.POST("/push", net.PushNotificationHandlerFor(e.grid))

	e.r.GET("/ws", net.WSHandlerFor(e.grid))
	e.l.Debug("gin endpoint setup complete")
}

func (e *Endpoint) Start() {
	e.l.Info("starting endpoint")
	err := e.r.Run()
	if err != nil {
		e.l.Fatal("could not run endpoint", zap.Error(err))
		return
	}
}
