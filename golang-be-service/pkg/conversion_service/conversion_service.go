package conversion_service

import (
	"strconv"

	resty "github.com/go-resty/resty/v2"
)

var API_KEY = "GbMiAb1Z8LCO5ETYixBrNvBSJhFXB57Y"
var BASE_URL = "https://api.apilayer.com/fixer/convert"

type ConversionService interface {
	Convert(amount float64, from string, to string) (float64, error)
}

type Converter struct {
	err error
}

func (c *Converter) Convert(amount float64, from string, to string) (float64, error) {
	client := resty.New()
	resp, err := client.R().SetHeader("apikey", API_KEY).Get(BASE_URL + "?from=" + from + "&to=" + to + "&amount="+ strconv.FormatFloat(amount, 'f', -1, 64))
	if err != nil {
		c.err = err
		return -1, err
	}
	return resp.Result().(map[string]interface{})["result"].(float64), nil
}
