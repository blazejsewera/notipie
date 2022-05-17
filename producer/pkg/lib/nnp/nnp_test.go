package nnp_test

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducer(t *testing.T) {
	t.Run("pushes notification", func(t *testing.T) {
		// given
		expected := testNotification
		mockSrv := httptest.NewServer(pushNotificationHandlerFor(t, expected))
		defer mockSrv.Close()

		producer := nnp.New(mockSrv.URL)

		// when
		appID := producer.Push(testNotification)

		// then
		assert.Equal(t, expected.AppID, appID)
	})
}

func pushNotificationHandlerFor(t testing.TB, expected model.AppNotification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// when
		an := deserializeAppNotification(t, r)

		// then
		assert.Equal(t, expected, an)

		_, err := w.Write([]byte(an.AppID))
		if err != nil {
			t.Fatal("write response: ", err)
		}
	}
}

func deserializeAppNotification(t testing.TB, r *http.Request) model.AppNotification {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal("read body: ", err)
	}
	anj := string(body)
	an, err := model.AppNotificationFromJSON(anj)
	if err != nil {
		t.Fatal("deserialize app notification json: ", err)
	}

	return an
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
