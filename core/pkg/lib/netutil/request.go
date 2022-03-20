package netutil

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func PostReq(c *http.Client, url string, contentType string, body string) (responseBody string, err error) {
	bodyReader := strings.NewReader(body)
	res, err := c.Post(url, contentType, bodyReader)
	if err != nil {
		return "", fmt.Errorf("post: %s", err)
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("post: read all bytes from response body: %s", err)
	}
	return string(resBytes), nil
}

func GetReq(c *http.Client, url string) (responseBody string, err error) {
	res, err := c.Get(url)
	if err != nil {
		return "", fmt.Errorf("get: %s", err)
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("get: read all bytes from response body: %s", err)
	}
	return string(resBytes), nil
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

func (c *WSReaderClient) Connect(url string) (err error) {
	c.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
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
