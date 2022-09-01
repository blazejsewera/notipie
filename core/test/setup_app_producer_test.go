package test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
)

type appProducer struct {
	appID   string
	t       testing.TB
	pushURL url.URL
	cl      *http.Client
}

func newAppProducer(t testing.TB, port int) *appProducer {
	return &appProducer{
		t:       t,
		pushURL: pushURL(port),
		cl:      &http.Client{Timeout: 200 * time.Millisecond},
	}
}

func (c *appProducer) pushNotification(notification model.AppNotification) {
	c.t.Helper()

	notificationJSON, err := notification.ToJSON()
	if err != nil {
		c.t.Fatal(err)
	}
	status, res, err := netutil.PostReq(c.cl, c.pushURL, "application/json", notificationJSON)
	if err != nil {
		c.t.Fatal(err)
	}

	assertStatusCreated(c.t, status)

	appIDRes, err := model.AppIDResponseFromReader(res)
	if err != nil {
		c.t.Fatal(err)
	}

	c.appID = appIDRes.AppID
	c.t.Log("appRestClient: pushNotification: success\tappID:", c.appID)
}

func assertStatusCreated(t testing.TB, statusCode int) {
	assert.Equal(t, http.StatusCreated, statusCode)
}
