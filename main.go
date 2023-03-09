package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
)

type Observation struct {
	XMLName xml.Name `xml:"current_observation"`
	Location string `xml:"location"`
	StationId string `xml:"station_id"`
	Latitute float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	ObservationTime string `xml:"observation_time"`
	ObservationTimeRfc822 string `xml:"observation_time_rfc822"`
	Weather string `xml:"weather"`
	TemperatureString string `xml:"temperature_string"`
	TempF float64 `xml:"temp_f"`
	TempC float64 `xml:"temp_c"`
	RelativeHumidity int `xml:"relative_humidity"`
	WindString string `xml:"wind_string"`
	WindDir string `xml:"wind_dir"`
	WindDegrees int `xml:"wind_degrees"`
	WindMph float64 `xml:"wind_mph"`
	WindKt int `xml:"wind_kt"`
	PressureString string `xml:"pressure_string"`
	PressureMb float64 `xml:"pressure_mb"`
	PressureIn float64 `xml:"pressure_in"`
	DewpointString string `xml:"dewpoint_string"`
	DewpointF float64 `xml:"dewpoint_f"`
	DewpointC float64 `xml:"dewpoint_c"`
	WindchillString string `xml:"windchill_string"`
	WindchillF float64 `xml:"windchill_f"`
	WindchillC float64 `xml:"windchill_c"`
	VisibilityMi float64 `xml:"visibility_mi"`
}

func main() {
	resp, err := http.Get("https://w1.weather.gov/xml/current_obs/KPDX.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	resp.Body.Close()

	var observation Observation
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&observation)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("%#v \n", observation)
}
