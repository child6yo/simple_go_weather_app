package main

import (
	"encoding/json"
)

type Responces interface {
	ParceResponce(data []byte) error
}

type WeatherAPIResponce struct {
	Result Weather `json:"current"`
}

type Weather struct {
	Weather float64 `json:"temperature_2m"`
}

type CoordinatesAPIResponce struct {
	Results []Coordinates `json:"results"`
}

type Coordinates struct {
	City string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (w *WeatherAPIResponce)ParceResponce(data []byte) error {
	err := json.Unmarshal(data, &w)
	if err != nil {
		return ParceError
	}
	return nil
}

func (c *CoordinatesAPIResponce)ParceResponce(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil || c.Results == nil {
		return ParceError
	}
	return nil
}

func ParceAnyResponce(r Responces, data []byte) error {
	err := r.ParceResponce(data)
	return err
}