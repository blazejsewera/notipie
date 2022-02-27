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
	c.initGin(config.prod)
	c.initLogger(config.prod)
	c.initGrid()
	c.initEndpoint()
}

func (c *AppContext) initGin(prod bool) {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func (c *AppContext) initLogger(prod bool) {
	log.Init(prod)
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
	c.ep.Run()
}
