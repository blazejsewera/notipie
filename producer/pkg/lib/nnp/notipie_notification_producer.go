package nnp

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"net/http"
	"net/url"
	"time"
)

type Producer interface {
	Push(notification model.AppNotification) error
}

type ProducerImpl struct {
	URL   url.URL
	c     *http.Client
	appID string
}

func NewProducer(rawURL, appID string) Producer {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		panic(fmt.Sprint("parse url:", err))
	}
	return &ProducerImpl{
		URL:   *parsedUrl,
		c:     http.DefaultClient,
		appID: appID,
	}
}

func (p *ProducerImpl) Push(notification model.AppNotification) error {
	notification.AppID = p.appID
	if notification.Timestamp == "" {
		notification.Timestamp = time.Now().Format(timeformat.RFC3339Milli)
	}

	notificationJSON, err := notification.ToJSON()
	if err != nil {
		return err
	}

	status, resBody, err := netutil.PostReq(p.c, p.URL, "application/json", notificationJSON)
	if err != nil {
		return err
	}

	if status != http.StatusCreated {
		return fmt.Errorf("push notification: server did not respond with correct status, status: %d", status)
	}

	appIDRes, err := model.AppIDResponseFromReader(resBody)
	if err != nil {
		return fmt.Errorf("push notification: %s", err)
	}

	p.appID = appIDRes.AppID
	return nil
}
