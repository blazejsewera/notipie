package net

import (
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var l = log.For("impl").Named("net").Named("server")

func PingHandler(c *gin.Context) {
	logRequest(l, c, "/")
	c.String(http.StatusOK, "OK")
}

func PreflightHandler(c *gin.Context) {
	logRequest(l, c, "preflight")
	c.Header("Access-Control-Allow-Origin", "*") // TODO: replace this dev value
	c.Header("Access-Control-Request-Method", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Connection", "keep-alive")
}

func PushNotificationHandlerFor(grid grid.Grid) gin.HandlerFunc {
	return func(c *gin.Context) {
		logRequest(l, c, "push")
		notification := model.AppNotification{}
		err := c.ShouldBindJSON(&notification)
		if err != nil {
			l.Error("error when binding json", zap.Error(err))
			return
		}
		l.Debug("received notification", zap.Reflect("notification", notification))
		appID := grid.ReceiveAppNotification(notification)
		c.JSON(http.StatusCreated, gin.H{"appId": appID})
	}
}

func GetNotificationsHandlerFor(grid grid.Grid) gin.HandlerFunc {
	return func(c *gin.Context) {
		logRequest(l, c, "notifications")
		username := c.DefaultQuery("user", "root") // TODO: add user auth
		userProxy, err := grid.GetUserProxy(username)
		if err != nil {
			l.Error("could not get user proxy", zap.Error(err))
			return
		}
		c.JSON(http.StatusOK, gin.H{"notifications": userProxy.GetAllNotifications()})
	}
}

func WSHandlerFor(grid grid.Grid) gin.HandlerFunc {
	upgrader := createUpgrader()

	return func(c *gin.Context) {
		logRequest(l, c, "ws")
		username := c.DefaultQuery("user", "root") // TODO: add user auth
		userProxy, err := grid.GetUserProxy(username)
		if err != nil {
			l.Error("could not get user proxy", zap.Error(err))
			return
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			l.Error("could not upgrade conn", zap.Error(err))
			return
		}
		userProxy.RegisterClient(conn)
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
