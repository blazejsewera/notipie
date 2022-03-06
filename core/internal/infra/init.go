package infra

import (
	"github.com/blazejsewera/notipie/core/internal/impl"
	"github.com/blazejsewera/notipie/core/internal/impl/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	gr grid.Grid
	ap grid.AppProxy
	up grid.UserProxy
	ep *impl.Endpoint
}

func (c *AppContext) Init(config Config) {
	c.initLogger(config.prod)
	c.initGin(config.prod)
	c.initGrid()
	c.initEndpoint()
}

func (c *AppContext) initLogger(prod bool) {
	log.Init(prod)
}

func (c *AppContext) initGin(prod bool) {
	l := log.For("infra").Named("init")
	if prod {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = voidWriter{}
		l.Info("gin is running in Debug mode")
	}
}

type voidWriter struct{}

func (v voidWriter) Write([]byte) (n int, err error) {
	return 0, nil
}

func (c *AppContext) initGrid() {
	c.gr = grid.NewGrid(ws.DefaultClientHubFactory{})
	c.gr.Start()
}

func (c *AppContext) initEndpoint() {
	c.ep = impl.NewEndpoint(c.gr)
	c.ep.Setup()
}

func (c *AppContext) Start() {
	c.ep.Start()
}
