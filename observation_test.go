package main

import (
	"regexp"
	"testing"
)

func TestGetWeatherValidLocationKPDX(t *testing.T) {
	location := "KPDX"
	want := regexp.MustCompile(`"location":"Portland, Portland International Airport, OR"`)
	js, err := GetWeather(location)
	if !want.MatchString(js) || err != nil {
		t.Fatalf("GetWeather(%q) = %q, want match for %q", location, js, want)
	}
}

func TestGetWeatherValidLocationKSEA(t *testing.T) {
	location := "KSEA"
	want := regexp.MustCompile(`"location":"Seattle, Seattle-Tacoma International Airport, WA"`)
	js, err := GetWeather(location)
	if !want.MatchString(js) || err != nil {
		t.Fatalf("GetWeather(%q) = %q, want match for %q", location, js, want)
	}
}

func TestNoLocation(t *testing.T) {
	js, err := GetWeather("")
	if js != "" || err == nil {
		t.Fatalf("GetWeather(%q) = %q, want error", "", js)
	}
}

func TestInvalidLocation(t *testing.T) {
	location := "invalid"
	js, err := GetWeather(location)
	if js != "" || err == nil {
		t.Fatalf("GetWeather(%q) = %q, want error", location, js)
	}
}
