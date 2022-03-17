package impl

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/net"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Endpoint struct {
	cfg  EndpointConfig
	r    *gin.Engine
	grid grid.Grid
	l    *zap.Logger
}

type EndpointConfig struct {
	Address   string `json:"address"`
	Port      int    `json:"port"`
	Prefix    string `json:"prefix"`
	Root      string `json:"root"`
	Push      string `json:"push"`
	WebSocket string `json:"webSocket"`
}

func NewEndpoint(endpointConfig EndpointConfig, grid grid.Grid) *Endpoint {
	return &Endpoint{
		cfg:  endpointConfig,
		r:    gin.New(),
		grid: grid,
		l:    log.For("impl").Named("endpoint"),
	}
}

func (e *Endpoint) Setup() {
	e.r.Use(gin.Recovery())

	root := e.cfg.Prefix + e.cfg.Root
	e.r.GET(root, net.PingHandler)

	push := e.cfg.Prefix + e.cfg.Push
	e.r.OPTIONS(push, net.PreflightHandler)
	e.r.POST(push, net.PushNotificationHandlerFor(e.grid))

	ws := e.cfg.Prefix + e.cfg.WebSocket
	e.r.GET(ws, net.WSHandlerFor(e.grid))
	e.l.Debug("gin endpoint setup complete")
}

func (e *Endpoint) Start() {
	addr := fmt.Sprintf("%s:%d", e.cfg.Address, e.cfg.Port)
	e.l.Info("starting endpoint", zap.String("addr", addr))
	err := e.r.Run(addr)
	if err != nil {
		e.l.Fatal("could not run endpoint", zap.Error(err))
		return
	}
}
