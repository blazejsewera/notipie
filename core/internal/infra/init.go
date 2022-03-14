package infra

import (
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/internal/impl"
	"github.com/blazejsewera/notipie/core/internal/impl/broadcast"
	"github.com/blazejsewera/notipie/core/internal/impl/persistence"
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
	c.initLogger(config.Prod)
	c.initGin(config.Prod)
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
	r := persistence.MemRepositoryFactory
	b := broadcast.WebSocketBroadcasterFactory
	c.gr = grid.NewGrid(r, b)
	c.gr.Start()
}

func (c *AppContext) initEndpoint() {
	c.ep = impl.NewEndpoint(c.gr)
	c.ep.Setup()
}

func (c *AppContext) Start() {
	c.ep.Start()
}
