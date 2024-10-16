package main

import (
	"fmt"
	"unicode"
	"log"
)

type Error string

func (e Error) Error() string { return string(e) }



const (
	InputError Error = "error: incorrect input"
	SymbolsError Error = "error: incorrect symbols in input (only latin are accepted)"
	RequestError Error = "error: cannot get responce"
	ReadError Error = "error: cannot read data"
	ParseError Error = "error: cannot parce data"
	
	WeatherAPIUrl = "https://api.open-meteo.com/v1/forecast?latitude=%g&longitude=%g&current=temperature_2m"
	CoordinatesAPIUrl = "https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json"
)


func CityInput() (string, error) {
	var city string
	_, err := fmt.Scan(&city)
	if err != nil { 
		return "", InputError
	}
	return city, nil
}

func InputCheck(city string) error {
	for _, l := range city {
		if !(unicode.In(l, unicode.Latin)) {
			return SymbolsError
		}
	}
	return nil
}

func CityHandler(city string) error {
	var weatherdata WeatherAPIResponse
	var coordinatesdata CoordinatesAPIResponse

	fcoordinatesurl := fmt.Sprintf(CoordinatesAPIUrl, city)
	coordparcer := Request{fcoordinatesurl}
	rawcdata, err := coordparcer.GetResponse()
	if err != nil {
		return err
	}
	err = ParceAnyResponse(&coordinatesdata, rawcdata)
	if err != nil {
		return err
	}
	cdata := coordinatesdata.Results[0]

	fweatherurl := fmt.Sprintf(WeatherAPIUrl, cdata.Latitude, cdata.Longitude)
	weatherparcer := Request{fweatherurl}
	rawwdata, err := weatherparcer.GetResponse()
	if err != nil {
		return err
	}
	err = ParceAnyResponse(&weatherdata, rawwdata)
	if err != nil {
		return err
	}
	wdata := weatherdata.Result.Weather

	fmt.Printf("Сейчас в городе %s %.f °C", cdata.City, wdata)

	return nil
}

func main() {
	city, err := CityInput()
	if err != nil {
		log.Fatal(err)
	}

	err = InputCheck(city)
	if err != nil {
		log.Fatal(err)
	}

	err = CityHandler(city)
	if err != nil {
		log.Fatal(err)
	}
}
