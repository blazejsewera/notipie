package ws

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
	"io"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	UUID           string
	hub            ClientHub
	conn           *websocket.Conn
	send           chan []byte
	l              *zap.Logger
	writeWait      time.Duration
	pongWait       time.Duration
	pingPeriod     time.Duration
	maxMessageSize int64
	newline        []byte
	space          []byte
}

func NewClient(uuid string, hub ClientHub, conn *websocket.Conn) *Client {
	writeWait := 10 * time.Second
	pongWait := 60 * time.Second
	pingPeriod := (pongWait * 9) / 10
	maxMessageSize := int64(8192)

	newline := []byte{'\n'}
	space := []byte{' '}

	return &Client{
		UUID:           uuid,
		hub:            hub,
		conn:           conn,
		send:           make(chan []byte, 256),
		l:              log.For("impl").Named("net").Named("ws"),
		writeWait:      writeWait,
		pongWait:       pongWait,
		pingPeriod:     pingPeriod,
		maxMessageSize: maxMessageSize,
		newline:        newline,
		space:          space,
	}
}

func (c *Client) readPump() {
	defer closeConnFor(c)

	err := c.setWSParams()
	if err != nil {
		return
	}
	c.readWholeMessage()
}

func closeConnFor(c *Client) {
	c.l.Debug("closing conn", zap.String("uuid", c.UUID))
	c.hub.GetUnregisterChan() <- c.UUID
	err := c.conn.Close()
	if err != nil {
		c.l.Error("could not close websocket", zap.Error(err))
		return
	}
}

func (c *Client) setWSParams() error {
	c.conn.SetReadLimit(c.maxMessageSize)
	errOrNil := c.setReadDeadline()
	c.conn.SetPongHandler(func(string) error {
		return c.setReadDeadline()
	})
	return errOrNil
}

func (c *Client) setReadDeadline() error {
	err := c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
	if err != nil {
		c.l.Error("could not set read deadline", zap.Error(err))
		return err
	}
	return nil
}

func (c *Client) readWholeMessage() {
	for {
		err := c.readMessage()
		if err != nil {
			c.l.Debug("websocket closed on the other end")
			return
		}
	}
}

func (c *Client) readMessage() error {
	n, notificationBytes, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
			c.l.Error("websocket unexpectedly closed")
		}
		return err
	}
	c.l.Debug(
		"read message from websocket",
		zap.String("uuid", c.UUID),
		zap.Int("bytesRead", n),
		zap.ByteString("notificationBytes", notificationBytes),
	)
	notificationBytes = bytes.TrimSpace(bytes.Replace(notificationBytes, c.newline, c.space, -1))
	c.l.Debug("trimmed whitespace", zap.ByteString("notificationBytes", notificationBytes))
	return nil
}

func (c *Client) writePump() {
	ticker := c.getTicker()
	defer c.stopTickerAndCloseConn(ticker)

	for {
		select {
		case message, ok := <-c.send:
			err := c.broadcastMessage(message, ok)
			if err != nil {
				return
			}
		case <-ticker.C:
			c.ping()
		}
	}
}

func (c *Client) getTicker() *time.Ticker {
	return time.NewTicker(c.pingPeriod)
}

func (c *Client) stopTickerAndCloseConn(ticker *time.Ticker) {
	ticker.Stop()
	_ = c.conn.Close()
	c.l.Debug("closed websocket connection", zap.String("uuid", c.UUID))
}

func (c *Client) broadcastMessage(message []byte, ok bool) error {
	err := c.setWriteDeadline()
	if err != nil {
		return err
	}
	if !ok {
		c.sendCloseMessage()
		return fmt.Errorf("sent close message")
	}

	c.writeMessage(message)
	c.l.Debug("broadcast message", zap.String("uuid", c.UUID), zap.ByteString("message", message))
	return nil
}

func (c *Client) setWriteDeadline() error {
	err := c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
	if err != nil {
		c.l.Error("could not set write deadline", zap.Error(err))
		return err
	}
	return nil
}

func (c *Client) sendCloseMessage() {
	_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
	c.l.Debug("sent close message", zap.String("uuid", c.UUID))
}

func (c *Client) writeMessage(message []byte) {
	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}

	c.write(w, message)
	c.handleQueuedMessages(w)

	c.close(w)
}

func (c *Client) write(w io.Writer, message []byte) {
	_, err := w.Write(message)
	if err != nil {
		c.l.Error("error when writing message to websocket", zap.Error(err))
	}
}

func (c *Client) handleQueuedMessages(w io.Writer) {
	n := len(c.send)
	c.l.Debug("handling queued messages", zap.Int("unhandledMessages", n))
	for i := 0; i < n; i++ {
		c.write(w, c.newline)
		c.write(w, <-c.send)
	}
}

func (c *Client) close(w io.Closer) {
	err := w.Close()
	if err != nil {
		c.l.Error("error when closing writer", zap.Error(err))
		return
	}
	c.l.Debug("closed writer", zap.String("uuid", c.UUID))
}

func (c *Client) ping() {
	err := c.setWriteDeadline()
	if err != nil {
		return
	}
	if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
		c.l.Warn("could not ping websocket", zap.Error(err))
		return
	}
	c.l.Debug("pinged websocket", zap.String("uuid", c.UUID))
}
