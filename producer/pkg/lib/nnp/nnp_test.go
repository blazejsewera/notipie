package nnp_test

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducer(t *testing.T) {
	t.Run("pushes notification", func(t *testing.T) {
		// given
		expected := testNotification
		ms := newMockServer(t)
		defer ms.close()

		producer := nnp.NewProducer(ms.URL)

		// when
		appID, err := producer.Push(testNotification)

		// then
		if assert.NoError(t, err) {
			ms.validateRequest(expected)
			assert.Equal(t, expected.AppID, appID)
		}
	})
}

type mockServer struct {
	URL      string
	received model.AppNotification
	s        *httptest.Server
	t        testing.TB
}

func newMockServer(t testing.TB) *mockServer {
	m := &mockServer{t: t}
	m.s = httptest.NewServer(http.HandlerFunc(m.pushNotificationHandler))
	m.URL = m.s.URL
	return m
}

func (m *mockServer) validateRequest(expected model.AppNotification) {
	assert.Equal(m.t, expected, m.received, "server did not get the expected request")
}

func (m *mockServer) pushNotificationHandler(w http.ResponseWriter, r *http.Request) {
	m.received = m.deserializeAppNotification(r)

	w.WriteHeader(http.StatusCreated)
	appID, err := model.AppIDResponseOf(m.received.AppID).ToJSON()
	if err != nil {
		m.t.Fatal("marshal app id response:", err)
	}
	_, err = w.Write(appID)
	if err != nil {
		m.t.Fatal("write response:", err)
	}
}

func (m *mockServer) deserializeAppNotification(r *http.Request) model.AppNotification {
	an, err := model.AppNotificationFromReader(r.Body)
	if err != nil {
		m.t.Fatal("deserialize app notification json: ", err)
	}
	return an
}

func (m *mockServer) close() {
	m.s.Close()
}

var testNotification = model.AppNotification{
	HashableNetNotification: model.HashableNetNotification{
		AppName:    "AppName",
		AppID:      "AppID",
		AppImgURI:  "AppImgURI",
		Title:      "Title",
		Subtitle:   "Subtitle",
		Body:       "Body",
		ExtURI:     "ExtURI",
		ReadURI:    "ReadURI",
		ArchiveURI: "ArchiveURI",
	},
	ID:        "ID",
	Timestamp: "Timestamp",
	Read:      true,
	ApiKey:    "ApiKey",
}
