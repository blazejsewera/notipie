package test

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/infra"
	"github.com/blazejsewera/notipie/core/internal/model"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

type appRestClient struct {
	appID string
	addr  string
	t     testing.TB
	cl    *http.Client
}

func initRestClient(t testing.TB, config infra.Config) *appRestClient {
	cl := &http.Client{Timeout: 200 * time.Millisecond}
	addr := fmt.Sprintf(
		"http://%s:%d%s",
		config.EndpointConfig.Address,
		config.EndpointConfig.Port,
		config.EndpointConfig.Prefix,
	)
	return &appRestClient{
		addr: addr,
		t:    t,
		cl:   cl,
	}
}

func (c *appRestClient) pushNotification(notification model.AppNotification) {
	c.t.Helper()
	logPrefix := "pushNotification: "

	url := c.getURL(push)
	n := strings.NewReader(notification.ToJSON())
	res, err := c.cl.Post(url, "application/json", n)
	if err != nil {
		c.t.Fatal(logPrefix+"send request", err)
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.t.Fatal(logPrefix+"could not read response body", err)
	}

	c.appID = string(resBytes)
	c.t.Log(logPrefix+"successful, appID:", c.appID)
}

func (c *appRestClient) getNotifications() {

}

func (c *appRestClient) getURL(path string) string {
	return c.addr + path
}
