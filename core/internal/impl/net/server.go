package net

import (
	"github.com/blazejsewera/notipie/core/internal/impl/grid"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
	"net/http"

	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func PingHandler(c *gin.Context) {
	l := log.For("impl").Named("net").Named("server")
	logRequest(l, c, "/")
	c.String(http.StatusOK, "OK")
}

func PreflightHandler(c *gin.Context) {
	l := log.For("impl").Named("net").Named("server")
	logRequest(l, c, "preflight")
	c.Header("Access-Control-Allow-Origin", "*") // TODO: replace this dev value
	c.Header("Access-Control-Request-Method", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Connection", "keep-alive")
}

func PushNotificationHandlerFor(grid grid.Grid) gin.HandlerFunc {
	l := log.For("impl").Named("net").Named("server")

	return func(c *gin.Context) {
		logRequest(l, c, "push")
		notification := model.AppNotification{}
		err := c.ShouldBindJSON(&notification)
		if err != nil {
			l.Error("error when binding json", zap.Error(err))
			return
		}
		l.Debug("received notification", zap.Reflect("notification", notification))
		grid.GetAppNotificationChan() <- notification
	}
}

func WSHandlerFor(grid grid.Grid) gin.HandlerFunc {
	l := log.For("impl").Named("net").Named("server")
	upgrader := createUpgrader()

	return func(c *gin.Context) {
		logRequest(l, c, "ws")
		username := c.DefaultQuery("user", "root") // TODO: add user auth
		userProxy, err := grid.GetUserProxy(username)
		if err != nil {
			l.Error("could not get user proxy", zap.Error(err))
			return
		}
		hub := userProxy.GetClientHub()
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			l.Error("could not upgrade conn", zap.Error(err))
			return
		}
		hub.GetRegisterChan() <- conn
	}
}

func createUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO: replace this dev value
		},
	}
}

func logRequest(l *zap.Logger, c *gin.Context, name string) {
	l.Debug(name, zap.String("from", c.Request.RemoteAddr))
}
