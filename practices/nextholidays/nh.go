package nh

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type holiday struct {
	Date        string      `json:"date"`
	LocalName   string      `json:"localName"`
	Name        string      `json:"name"`
	CountryCode string      `json:"countryCode"`
	Counties    interface{} `json:"counties"`
	LaunchYear  interface{} `json:"launchYear"`
	Type        string      `json:"type"`
	Fixed       bool        `json:"fixed"`
	Global      bool        `json:"global"`
}

var nzHolidays []*holiday

const URL = "https://date.nager.at/api/v2/publicholidays/"
const ymdFmt = "20060102"
const jointStr = "/"
const sep = ""

func GetHolidays() (nHolidays map[string]*holiday, err error) {
	t := time.Now()
	year := t.Year()
	reqParams := []string{URL, strconv.Itoa(year), jointStr, "NZ"}
	reqURL := strings.Join(reqParams, sep)
	resp, _ := http.Get(reqURL)
	payload, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(payload, &nzHolidays)
	if err != nil {
		return nil, err
	}

	nHolidays = make(map[string]*holiday)
	currentDateStr := time.Now().Format(ymdFmt)

	for _, h := range nzHolidays {
		strKey := strings.Replace(h.Date, "-", "", 2)
		if currentDateStr <= strKey && h.Global {
			nHolidays[strKey] = h
		}
	}

	return
}
