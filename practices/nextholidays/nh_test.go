package main

import (
	"log"
	"strings"
	"testing"
	"time"
)

func Test_getHolidays(t *testing.T) {
	nHolidays, err := getHolidays()
	currentDateStr := time.Now().Format("20060102")
	if err != nil {
		log.Fatal(err)
	}

	for _, tt := range nHolidays {
		t.Run("Test next holiday gather than currentDate", func(t *testing.T) {
			strKey := strings.Replace(tt.Date, "-", "", 2)
			if strKey < currentDateStr {
				t.Errorf("getHolidays() error = %v, wantErr %v", err, tt.Date)
				return
			}
		})
	}
}
