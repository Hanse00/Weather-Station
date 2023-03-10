package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
)

type Observation struct {
	XMLName xml.Name `xml:"current_observation" json:"-"`
	Location string `xml:"location" json:"location"`
	StationId string `xml:"station_id" json:"station_id"`
	Latitute float64 `xml:"latitude" json:"latitude"`
	Longitude float64 `xml:"longitude" json:"longitude"`
	ObservationTime string `xml:"observation_time" json:"observation_time"`
	ObservationTimeRfc822 string `xml:"observation_time_rfc822" json:"observation_time_rfc822"`
	Weather string `xml:"weather" json:"weather"`
	TemperatureString string `xml:"temperature_string" json:"temperature_string"`
	TempF float64 `xml:"temp_f" json:"temp_f"`
	TempC float64 `xml:"temp_c" json:"temp_c"`
	RelativeHumidity int `xml:"relative_humidity" json:"relative_humidity"`
	WindString string `xml:"wind_string" json:"wind_string"`
	WindDir string `xml:"wind_dir" json:"wind_dir"`
	WindDegrees int `xml:"wind_degrees" json:"wind_degrees"`
	WindMph float64 `xml:"wind_mph" json:"wind_mph"`
	WindKt int `xml:"wind_kt" json:"wind_kt"`
	PressureString string `xml:"pressure_string" json:"pressure_string"`
	PressureMb float64 `xml:"pressure_mb" json:"pressure_mb"`
	PressureIn float64 `xml:"pressure_in" json:"pressure_in"`
	DewpointString string `xml:"dewpoint_string" json:"dewpoint_string"`
	DewpointF float64 `xml:"dewpoint_f" json:"dewpoint_f"`
	DewpointC float64 `xml:"dewpoint_c" json:"dewpoint_c"`
	WindchillString string `xml:"windchill_string" json:"windchill_string"`
	WindchillF float64 `xml:"windchill_f" json:"windchill_f"`
	WindchillC float64 `xml:"windchill_c" json:"windchill_c"`
	VisibilityMi float64 `xml:"visibility_mi" json:"visibility_mi"`
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

	js, err := json.Marshal(observation)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println(string(js))
}
