package test

import (
	"testing"

	resty "github.com/go-resty/resty/v2"
	assert "github.com/stretchr/testify/assert"
)

var (
	BASE_URL = "http://localhost:3001/api/v1"
)


func TestLogin(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","password":"admin"}`).Post(BASE_URL + "/login")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestSignUp(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","password":"admin","email":"admin123@gmail.com","firstname":"admin","lastname":"admin"}`).Post(BASE_URL + "/signup")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}


func TestBalance(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/balance")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}


func TestGetProfile(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/profile")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}


func TestUpdateProfile(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","password":"admin"}`).Put(BASE_URL + "/profile")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}


func TestTransfer(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"username":"admin","amount":"1000"}`).Post(BASE_URL + "/transfer")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}


func TestTransact(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"transaction_type":"admin","amount":"1000"}`).Post(BASE_URL + "/transact")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
