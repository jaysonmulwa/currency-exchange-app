package conversion_service

import (
	"strconv"

	resty "github.com/go-resty/resty/v2"
)

var API_KEY = "GbMiAb1Z8LCO5ETYixBrNvBSJhFXB57Y"

type ConversionService interface {
	Convert(amount int, from string, to string) (float64, error)
}

type Converter struct {
	err error
}

func (c *Converter) Convert(amount int, from string, to string) (float64, error) {
	client := resty.New()
	resp, err := client.R().SetHeader("apikey", API_KEY).Get("https://api.apilayer.com/fixer/convert?from=" + from + "&to=" + to + "&amount="+ strconv.Itoa(amount))
	if err != nil {
		c.err = err
		return -1, err
	}
	return resp.Result().(map[string]interface{})["result"].(float64), nil
}
