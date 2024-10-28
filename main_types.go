package main

import "encoding/xml"

type CountyItems struct {
	XMLName    xml.Name     `xml:"countyItems"`
	CountyLIst []CountyItem `xml:"countyItem"`
}

type CountyItem struct {
	CountyCode   string `xml:"countycode"`
	CountyName   string `xml:"countyname"`
	CountyCode01 string `xml:"countycode01"`
}

type TownItems struct {
	XMLName  xml.Name   `xml:"townItems"`
	TownList []TownItem `xml:"townItem"`
}

type TownItem struct {
	TownCode string `xml:"towncode"`
	TownName string `xml:"townname"`
}

type SelectItem struct {
	Country string
	Code    string
	Towns   []struct {
		Code string
		Town string
	}
}
