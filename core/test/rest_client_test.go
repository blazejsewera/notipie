package test

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"net/http"
	"testing"
	"time"
)

var addr = fmt.Sprintf(
	"http://%s:%d%s",
	config.EndpointConfig.Address,
	config.EndpointConfig.Port,
	config.EndpointConfig.Prefix,
)

var wsAddr = fmt.Sprintf(
	"ws://%s:%d%s",
	config.EndpointConfig.Address,
	config.EndpointConfig.Port,
	config.EndpointConfig.Prefix,
)

func getURL(addr string, path string) string {
	return addr + path
}

type appRestClient struct {
	appID string
	t     testing.TB
	cl    *http.Client
}

type appRes struct {
	AppID string `json:"appId"`
}

func newAppRestClient(t testing.TB) *appRestClient {
	cl := &http.Client{Timeout: 200 * time.Millisecond}
	return &appRestClient{
		t:  t,
		cl: cl,
	}
}

func (c *appRestClient) pushNotification(notification model.AppNotification) {
	c.t.Helper()
	logPrefix := "pushNotification: "

	url := getURL(addr, push)
	res, err := netutil.PostReq(c.cl, url, "application/json", notification.ToJSON())
	if err != nil {
		c.t.Fatal(err)
	}

	c.appID, err = appIdFromRes(res)
	if err != nil {
		c.t.Fatal(err)
	}

	c.log(logPrefix, "success: appID = "+c.appID)
}

func appIdFromRes(appResJSON string) (string, error) {
	a := appRes{}
	err := json.Unmarshal([]byte(appResJSON), &a)
	if err != nil {
		return "", fmt.Errorf("unmarshal app response: %s", err)
	}
	return a.AppID, nil
}

func (c *appRestClient) log(prefix string, message string) {
	c.t.Log("appRestClient:", prefix, message)
}

type userClient struct {
	notifications []model.ClientNotification
	t             testing.TB
}

type wsClient struct {
	userClient
	wsc   *netutil.WSReaderClient
	saved chan util.Signal
}

func newWSClient(t testing.TB) *wsClient {
	return &wsClient{userClient: userClient{t: t}, saved: make(chan util.Signal, 1)}
}

func (c *wsClient) connect() {
	url := getURL(wsAddr, ws)
	c.wsc = netutil.NewTestWSReaderClient(c.t)
	err := c.wsc.Connect(url)
	if err != nil {
		c.t.Fatal(err)
	}
	go c.readWS()
}

func (c *wsClient) readWS() {
	for {
		<-c.wsc.Saved
		lastLine := c.wsc.LineBuffer[len(c.wsc.LineBuffer)-1]
		notification, err := model.ClientNotificationFromJSON(lastLine)
		if err != nil {
			c.t.Error("ws client: unmarshal client notification:", err)
			continue
		}
		c.notifications = append(c.notifications, notification)
		c.saved <- util.Ping
	}
}

func (c *wsClient) close() {
	err := c.wsc.Close()
	if err != nil {
		c.t.Fatal(err)
	}
}

type userRestClient struct {
	userClient
}

func initUserRestClient(t testing.TB) *userRestClient {
	return &userRestClient{userClient: userClient{t: t}}
}

func (c *userRestClient) getNotifications() {
	c.notifications = append(c.notifications) // TODO: make request, append response
}
