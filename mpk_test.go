package mpk

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceFetch(t *testing.T) {
	service, err := New(&http.Client{})
	if assert.NoError(t, err) {
		_, err := service.GPS.Fetch(map[TransportationType][]string{
			TransportationTypeBus: []string{
				"a",
			},
		})

		assert.NoError(t, err)
	}
}
