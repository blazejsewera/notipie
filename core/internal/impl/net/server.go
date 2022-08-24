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

type Server struct {
	l *zap.Logger
}

func NewServer() *Server {
	return &Server{
		l: log.For("impl").Named("net").Named("server"),
	}
}

func (s *Server) PingHandler(c *gin.Context) {
	s.logRequest(c, "/")
	c.String(http.StatusOK, "OK")
}

func (s *Server) PreflightHandler(c *gin.Context) {
	s.logRequest(c, "preflight")
	c.Header("Access-Control-Allow-Origin", "*") // TODO: replace this dev value
	c.Header("Access-Control-Request-Method", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Connection", "keep-alive")
}

func (s *Server) PushNotificationHandlerFor(grid grid.Grid) gin.HandlerFunc {
	return func(c *gin.Context) {
		s.logRequest(c, "push")
		notification := model.AppNotification{}
		err := c.ShouldBindJSON(&notification)
		if err != nil {
			s.l.Error("error when binding json", zap.Error(err))
			return
		}
		s.l.Debug("received notification", zap.Reflect("notification", notification))
		appID := grid.ReceiveAppNotification(notification)
		c.JSON(http.StatusCreated, gin.H{"appId": appID})
	}
}

func (s *Server) GetNotificationsHandlerFor(grid grid.Grid) gin.HandlerFunc {
	return func(c *gin.Context) {
		s.logRequest(c, "notifications")
		username := c.DefaultQuery("user", "root") // TODO: add user auth
		userProxy, err := grid.GetUserProxy(username)
		if err != nil {
			s.l.Error("could not get user proxy", zap.Error(err))
			return
		}
		c.JSON(http.StatusOK, gin.H{"notifications": userProxy.GetAllNotifications()})
	}
}

func (s *Server) WSHandlerFor(grid grid.Grid) gin.HandlerFunc {
	upgrader := s.createUpgrader()

	return func(c *gin.Context) {
		s.logRequest(c, "ws")
		username := c.DefaultQuery("user", "root") // TODO: add user auth
		userProxy, err := grid.GetUserProxy(username)
		if err != nil {
			s.l.Error("could not get user proxy", zap.Error(err))
			return
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			s.l.Error("could not upgrade conn", zap.Error(err))
			return
		}
		userProxy.RegisterClient(conn)
	}
}

func (s *Server) createUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO: replace this dev value
		},
	}
}

func (s *Server) logRequest(c *gin.Context, name string) {
	s.l.Debug(name, zap.String("from", c.Request.RemoteAddr))
}
