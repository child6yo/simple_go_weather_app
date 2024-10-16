package main

import (
	"encoding/json"
)

type Responses interface {
	ParseResponse(data []byte) error
}

type WeatherAPIResponse struct {
	Result Weather `json:"current"`
}

type Weather struct {
	Weather float64 `json:"temperature_2m"`
}

type CoordinatesAPIResponse struct {
	Results []Coordinates `json:"results"`
}

type Coordinates struct {
	City string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (w *WeatherAPIResponse)ParseResponse(data []byte) error {
	err := json.Unmarshal(data, &w)
	if err != nil {
		return ParseError
	}
	return nil
}

func (c *CoordinatesAPIResponse)ParseResponse(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil || c.Results == nil {
		return ParseError
	}
	return nil
}

func ParceAnyResponse(r Responses, data []byte) error {
	err := r.ParseResponse(data)
	return err
}