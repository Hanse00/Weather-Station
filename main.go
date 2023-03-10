package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a location")
		os.Exit(1)
	}
	location := os.Args[1]

	js, err := GetWeather(location)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(js)
}
