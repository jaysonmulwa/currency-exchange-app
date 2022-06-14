package test

import (
	"testing"

	resty "github.com/go-resty/resty/v2"
	assert "github.com/stretchr/testify/assert"
)

//Auth Tests
func TestGetUserByCredentials(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","password":"admin"}`).Post(BASE_URL + "/login")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestCreateUser(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","password":"admin"}`).Post(BASE_URL + "/login")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
