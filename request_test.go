package main

import (
	"testing"
)

const minwantedbytes = 15

func TestGetResponce(t *testing.T) {
	cases := []struct {
		name string
		value Request
		want int
		err error
	}{
		{
			name: "geocoding",
			value: Request{"https://geocoding-api.open-meteo.com/v1/search?name=Moscow&count=1&language=ru&format=json"},
			want: minwantedbytes,
		},
		{
			name: "weather",
			value: Request{"https://api.open-meteo.com/v1/forecast?latitude=53.3606&longitude=83.7636&current=temperature_2m"},
			want: minwantedbytes,
		},
		{
			name: "empty url",
			value: Request{""},
			want: 0,
			err: RequestError,
		},
	}
	for _, tc := range cases  {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.GetResponce()
			if len(got) < tc.want || err != tc.err {
				t.Errorf("%q failed", tc.name)
			}
		})
	}
}