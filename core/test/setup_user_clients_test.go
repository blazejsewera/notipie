package test

import (
	"encoding/json"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type userClient struct {
	notifications []model.ClientNotification
	t             testing.TB
}

type userWSClient struct {
	userClient
	wsc   *netutil.WSReaderClient
	wsURL url.URL
	saved chan util.Signal
}

func newUserWSClient(t testing.TB, port int) *userWSClient {
	return &userWSClient{
		userClient: userClient{t: t},
		wsURL:      wsURL(port),
		saved:      make(chan util.Signal, 1),
	}
}

func (c *userWSClient) connect() {
	c.wsc = netutil.NewTestWSReaderClient(c.t)
	err := c.wsc.Connect(c.wsURL)
	if err != nil {
		c.t.Fatal(err)
	}
	go c.readWS()
}

func (c *userWSClient) readWS() {
	for {
		<-c.wsc.Saved
		lastLine := c.wsc.LineBuffer[len(c.wsc.LineBuffer)-1]
		notification, err := model.ClientNotificationFromJSON(lastLine)
		if err != nil {
			continue
		}
		c.notifications = append(c.notifications, notification)
		c.saved <- util.Ping
	}
}

func (c *userWSClient) close() {
	err := c.wsc.Close()
	if err != nil {
		c.t.Fatal(err)
	}
}

type userRestClient struct {
	userClient
	notificationsURL url.URL
	cl               *http.Client
}

type notificationsRes struct {
	Notifications []model.ClientNotification `json:"notifications"`
}

func notificationsFromRes(res []byte) ([]model.ClientNotification, error) {
	nRes := notificationsRes{}
	err := json.Unmarshal(res, &nRes)
	if err != nil {
		return nil, err
	}
	return nRes.Notifications, err
}

func newUserRestClient(t testing.TB, port int) *userRestClient {
	return &userRestClient{
		userClient:       userClient{t: t},
		notificationsURL: notificationsURL(port),
		cl:               &http.Client{Timeout: 200 * time.Millisecond},
	}
}

func (c *userRestClient) getNotifications() {
	status, res, err := netutil.GetReq(c.cl, c.notificationsURL)
	if err != nil {
		c.t.Fatal(err)
	}

	assertStatusOK(c.t, status)

	n, err := notificationsFromRes(res)
	if err != nil {
		c.t.Fatal(err)
	}

	c.notifications = append(c.notifications, n...)
}

func assertStatusOK(t testing.TB, statusCode int) {
	t.Helper()
	assert.Equal(t, http.StatusOK, statusCode)
}
