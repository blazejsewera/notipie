package test

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type appRestClient struct {
	appID string
	t     testing.TB
	cl    *http.Client
}

type appIDRes struct {
	AppID string `json:"appId"`
}

func newAppRestClient(t testing.TB) *appRestClient {
	return &appRestClient{
		t:  t,
		cl: &http.Client{Timeout: 200 * time.Millisecond},
	}
}

func (c *appRestClient) pushNotification(notification model.AppNotification) {
	c.t.Helper()

	status, res, err := netutil.PostReq(c.cl, pushURL, "application/json", notification.ToJSON())
	if err != nil {
		c.t.Fatal(err)
	}

	assertStatusCreated(c.t, status)

	c.appID, err = appIdFromRes(res)
	if err != nil {
		c.t.Fatal(err)
	}

	c.t.Log("appRestClient: pushNotification: success\t" + string(res))
}

func assertStatusCreated(t testing.TB, statusCode int) {
	assert.Equal(t, http.StatusCreated, statusCode)
}

func appIdFromRes(appResJSON []byte) (string, error) {
	a := appIDRes{}
	err := json.Unmarshal(appResJSON, &a)
	if err != nil {
		return "", fmt.Errorf("unmarshal app response: %s", err)
	}
	return a.AppID, nil
}
