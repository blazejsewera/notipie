package impl

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/api"

	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/net"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Endpoint struct {
	cfg    EndpointConfig
	r      *gin.Engine
	grid   grid.Grid
	server *net.Server
	l      *zap.Logger
}

type EndpointConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func NewEndpoint(endpointConfig EndpointConfig, grid grid.Grid) *Endpoint {
	return &Endpoint{
		cfg:    endpointConfig,
		r:      gin.New(),
		grid:   grid,
		server: net.NewServer(),
		l:      log.For("impl").Named("endpoint"),
	}
}

func (e *Endpoint) Setup() {
	e.r.Use(gin.Recovery())

	root := api.GetPath(api.Root)
	e.r.GET(root, e.server.PingHandler)

	push := api.GetPath(api.Push)
	e.r.OPTIONS(push, e.server.PreflightHandler)
	e.r.POST(push, e.server.PushNotificationHandlerFor(e.grid))

	notifications := api.GetPath(api.Notifications)
	e.r.GET(notifications, e.server.GetNotificationsHandlerFor(e.grid))

	ws := api.GetPath(api.WebSocket)
	e.r.GET(ws, e.server.WSHandlerFor(e.grid))
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
