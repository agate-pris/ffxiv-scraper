package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeObj := time.Unix(1560000000, 0)
	if timeObj.Year() != 2019 {
		t.Error(timeObj.Year())
	}
	if timeObj.Month() != 6 {
		t.Error(timeObj.Month())
	}
	if timeObj.Day() != 8 {
		t.Error(timeObj.Day())
	}
	if timeObj.Hour() != 22 {
		t.Error(timeObj.Hour())
	}
	if timeObj.Minute() != 20 {
		t.Error(timeObj.Hour())
	}
	if timeObj.Second() != 0 {
		t.Error(timeObj.Second())
	}
	location, e := time.LoadLocation("Asia/Tokyo")
	if e != nil {
		t.Fatal(e)
	}
	if timeObj.In(location).String() != "2019-06-08 22:20:00 +0900 JST" {
		t.Error(timeObj.In(location).String())
	}
}
