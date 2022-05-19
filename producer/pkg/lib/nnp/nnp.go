package nnp

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"io"
	"net/http"
	"net/url"
)

type Producer interface {
	Push(notification model.AppNotification) (appID string, err error)
}

type ProducerImpl struct {
	URL   url.URL
	c     *http.Client
	appID string
}

func NewProducer(rawURL string) *ProducerImpl {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		panic(fmt.Sprint("parse url:", err))
	}
	return &ProducerImpl{
		URL: *parsedUrl,
		c:   http.DefaultClient,
	}
}

func (p *ProducerImpl) Push(notification model.AppNotification) (appID string, err error) {
	status, resBody, err := netutil.PostReq(p.c, p.URL, "application/json", notification.ToJSON())
	if err != nil {
		return "", err
	}
	if status != http.StatusCreated {
		return "", fmt.Errorf("push notification: server did not respond with correct status, status: %d", status)
	}

	p.appID, err = appIDFromRes(resBody)

	if err != nil {
		return "", fmt.Errorf("push notification: %s", err)
	}

	return p.appID, nil
}

func appIDFromRes(r io.Reader) (string, error) {
	a := AppIDRes{}
	d := json.NewDecoder(r)
	err := d.Decode(&a)
	if err != nil {
		return "", fmt.Errorf("unmarshaling appID: %s", err)
	}
	return a.AppID, nil
}

type AppIDRes struct {
	AppID string `json:"appId"`
} // TODO: refactor it so that it is in core/model
