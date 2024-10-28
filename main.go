package main

import (
	"encoding/xml"
	"fmt"

	mlog "github.com/mike504110403/goutils/log"

	"github.com/valyala/fasthttp"
)

func main() {
	items := []SelectItem{}
	if countries, err := getCountry(); err != nil {
		mlog.Error(err.Error())
	} else {
		for _, country := range countries.CountyLIst {
			item := SelectItem{}
			if towns, err := getTown(country.CountyCode); err != nil {
				mlog.Error(err.Error())
			} else {
				item.Country = country.CountyName
				item.Code = country.CountyCode
				for _, town := range towns.TownList {
					item.Towns = append(item.Towns, struct {
						Code string
						Town string
					}{
						Code: town.TownCode,
						Town: town.TownName,
					})
				}
				items = append(items, item)
			}
		}
	}
	mlog.Info(fmt.Sprintf("%v", items))
}

func getCountry() (CountyItems, error) {
	res := CountyItems{}
	if body, err := sendGet("https://api.nlsc.gov.tw/other/ListCounty"); err != nil {
		return res, err
	} else {
		if err := xml.Unmarshal(body, &res); err != nil {
			return res, err
		}
	}
	return res, nil
}

func getTown(country string) (TownItems, error) {
	res := TownItems{}
	if body, err := sendGet("https://api.nlsc.gov.tw/other/ListTown/" + country); err != nil {
		return res, err
	} else {
		if err := xml.Unmarshal(body, &res); err != nil {
			return res, err
		}
	}
	return res, nil
}

func sendGet(url string) ([]byte, error) {
	client := &fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		return []byte{}, err
	} else {
		return resp.Body(), nil

	}
}
