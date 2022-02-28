package resource

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"service/cache"
)

type ResponseData struct {
	Uuid         string  `json:"uuid"`
	Komoditas    string  `json:"komoditas"`
	AreaProvinsi string  `json:"area_provinsi"`
	AreaKota     string  `json:"area_kota"`
	Size         string  `json:"size"`
	Price        string  `json:"price"`
	PriceUsd     float64 `json:"price_usd"`
	TglParsed    string  `json:"tgl_parsed"`
	Timestamp    string  `json:"timestamp"`
}

type CurrencyCount struct {
	Count int `json:"count"`
}

type Currency struct {
	Query   CurrencyCount          `json:"query"`
	Results map[string]interface{} `json:"results"`
}

func newResponseData() *ResponseData {
	return &ResponseData{}
}

func newCurrency() *Currency {
	return &Currency{}
}

func fetchMainResource() ([]*ResponseData, error) {
	req, err := http.NewRequest(
		"GET",
		"https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list",
		nil,
	)

	if err != nil {
		return nil, err
	}

	// Set 'Content-Type' header to 'application/json'.
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	list := make([]*ResponseData, 0)

	err = json.Unmarshal(buf, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func fetchCurrency() (*Currency, error) {
	req, err := http.NewRequest(
		"GET",
		"https://free.currconv.com/api/v7/convert?q=IDR_USD&apiKey=6c686cb7c5ac91832327",
		nil,
	)

	if err != nil {
		return nil, err
	}

	// Set 'Content-Type' header to 'application/json'.
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	c := newCurrency()

	err = json.Unmarshal(buf, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func Fetch() ([]*ResponseData, error) {
	list, err := fetchMainResource()

	if err != nil {
		return nil, err
	}

	c, err := cache.CurrencyCacheFactory()

	if err != nil {
		return nil, err
	}

	r := c.All()

	var t *Currency

	if _, ok := r.Data["IDR_USD"]; !ok {
		t, err = fetchCurrency()

		if err != nil {
			return nil, err
		}

		c.Put("IDR_USD", t.Results["IDR_USD"])
	}

	if t != nil {
		r.Data["IDR_USD"] = t.Results["IDR_USD"]
	}

	concrete, ok := r.Data["IDR_USD"].(map[string]interface{})

	if !ok {
		return nil, errors.New("Cannot infer type of r.Data['IDR_USD'].")
	}

	for i := 0; i < len(list); i++ {
		if list[i].Price != "" {
			price, err := strconv.ParseInt(list[i].Price, 10, 32)

			if err != nil {
				return nil, err
			}

			list[i].PriceUsd = float64(price) * concrete["val"].(float64)
		}
	}

	return list, nil
}
