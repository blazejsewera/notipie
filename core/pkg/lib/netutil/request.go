package netutil

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func PostReq(c *http.Client, url url.URL, contentType, body string) (
	statusCode int,
	responseBody io.Reader,
	err error,
) {
	bodyReader := strings.NewReader(body)
	res, err := c.Post(url.String(), contentType, bodyReader)
	statusCode = res.StatusCode

	if err != nil {
		err = fmt.Errorf("post: %s", err)
		return
	}

	responseBody = res.Body
	return
}

func GetReq(c *http.Client, url url.URL) (statusCode int, responseBody io.Reader, err error) {
	res, err := c.Get(url.String())
	statusCode = res.StatusCode

	if err != nil {
		err = fmt.Errorf("get: %s", err)
		return
	}

	responseBody = res.Body
	return
}

type WSReaderClient struct {
	Buffer     []byte
	LineBuffer []string
	Saved      chan util.Signal
	conn       *websocket.Conn
	t          testing.TB
	l          *zap.Logger
}

func NewTestWSReaderClient(t testing.TB) *WSReaderClient {
	return &WSReaderClient{Saved: make(chan util.Signal, 1), t: t}
}

func NewWSReaderClient() *WSReaderClient {
	l := log.For("netutil").Named("request").Named("wsReaderClient")
	return &WSReaderClient{Saved: make(chan util.Signal, 1), l: l}
}

func (c *WSReaderClient) Connect(url url.URL) (err error) {
	c.conn, _, err = websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return fmt.Errorf("ws client: connect: %s", err)
	}
	go c.readWS()
	return nil
}

func (c *WSReaderClient) readWS() {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		c.Buffer = append(c.Buffer, msg...)
		c.LineBuffer = append(c.LineBuffer, string(msg))
		c.Saved <- util.Ping
	}
}

func (c *WSReaderClient) writeErr(msg string, err error) {
	if c.t != nil {
		c.t.Helper()
		c.t.Log("ws client:", msg+":", err)
	}
	if c.l != nil {
		c.l.Debug(msg, zap.Error(err))
	}
}

func (c *WSReaderClient) Close() (err error) {
	_ = c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
	err = c.conn.Close()
	return
}
