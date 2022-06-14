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
	// given
	testNotification := model.ExampleAppNotification
	testNotification.AppID = ""
	testNotificationWithoutTimestamp := testNotification
	testNotificationWithoutTimestamp.Timestamp = ""

	t.Run("pushes notification the first time", func(t *testing.T) {
		// given
		ms := newMockServer(t)
		defer ms.close()

		producer := nnp.NewProducer(ms.URL, "")

		// when
		err := producer.Push(testNotification)

		// then
		if assert.NoError(t, err) {
			ms.validateRequest(testNotification)
		}
	})

	t.Run("pushes notification twice", func(t *testing.T) {
		// given
		ms := newMockServer(t)
		defer ms.close()

		producer := nnp.NewProducer(ms.URL, "")

		// when
		err := producer.Push(testNotification)

		// then
		assert.NoError(t, err)
		err = producer.Push(testNotification)
		if assert.NoError(t, err) {
			ms.validateSecondRequest(testNotification)
		}
	})

	t.Run("adds timestamp on push", func(t *testing.T) {
		// given
		ms := newMockServer(t)
		defer ms.close()

		producer := nnp.NewProducer(ms.URL, "")

		// when
		err := producer.Push(testNotificationWithoutTimestamp)

		// then
		if assert.NoError(t, err) {
			ms.validateRequestHasTimestamp()
		}
	})
}

type mockServer struct {
	URL      string
	received model.AppNotification
	appID    string
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

func (m *mockServer) validateSecondRequest(expected model.AppNotification) {
	expected.AppID = m.appID
	assert.Equal(m.t, expected, m.received, "server did not get the expected request, check appID")
}

func (m *mockServer) validateRequestHasTimestamp() {
	assert.NotEqual(m.t, "", m.received.Timestamp, "timestamp was not appended for app notification")
}

func (m *mockServer) pushNotificationHandler(w http.ResponseWriter, r *http.Request) {
	m.received = m.deserializeAppNotification(r)
	m.generateNewAppIDIfDoesNotExist()

	w.WriteHeader(http.StatusCreated)
	appID, err := model.AppIDResponseOf(m.appID).ToJSON()
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
		m.t.Fatal("deserialize app notification json:", err)
	}
	return an
}

func (m *mockServer) generateNewAppIDIfDoesNotExist() {
	if m.received.AppID != "" {
		return
	}
	m.appID = "AppID"
}

func (m *mockServer) close() {
	m.s.Close()
}
