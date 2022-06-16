package nnp

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"net/http"
	"time"
)

type Producer interface {
	Push(notification model.AppNotification) error
}

type ProducerImpl struct {
	c          *http.Client
	cfg        ProducerConfig
	appIDSaver AppIDSaver
}

func NewProducer(cfg ProducerConfig, appIDSaver AppIDSaver) Producer {
	return &ProducerImpl{
		c:          &http.Client{},
		cfg:        cfg,
		appIDSaver: appIDSaver,
	}
}

func (p *ProducerImpl) Push(notification model.AppNotification) error {
	notification.AppID = p.cfg.AppID
	if notification.Timestamp == "" {
		notification.Timestamp = time.Now().Format(timeformat.RFC3339Milli)
	}

	notificationJSON, err := notification.ToJSON()
	if err != nil {
		return err
	}

	status, resBody, err := netutil.PostReq(p.c, p.cfg.Endpoint.PushURL, "application/json", notificationJSON)
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

	p.cfg.AppID = appIDRes.AppID
	err = p.appIDSaver.SaveAppID(appIDRes.AppID)
	if err != nil {
		return err
	}

	return nil
}
