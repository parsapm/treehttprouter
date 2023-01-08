package treehttprouter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := newRoute()
	assert.NotNilf(t, r, "return nil router")
}

func TestAddHandler(t *testing.T) {
	r := newRoute()

	var h1 Handler = func(r *http.Request) error { return nil }
	var h2 Handler = func(r *http.Request) error { return nil }

	_ = r.addHandler("GET", &h1)

	_ = r.addHandler("POST", &h2)

	tr := &Route{
		handler: map[string]*Handler{
			"GET":  &h1,
			"POST": &h2,
		},
	}

	assert.Equal(t, tr, r)
}
