package main

import "testing"

func TestCityHandler(t *testing.T) {
	cases := []struct {
		name string
		value string
		err error
	}{
		{
			name: "Barnaul",
			value: "Barnaul",
			err: nil,
		},
		{
			name: "Empty",
			value: "",
			err: ParseError,
		},
	}
	for _, tc := range cases  {
		t.Run(tc.name, func(t *testing.T) {
			err := CityHandler(tc.value)
			if err != tc.err {
				t.Errorf("%q failed, get: %q", tc.name, err)
			}
		})
	}
}

func TestInputCheck(t *testing.T){
	cases := []struct {
		name string
		value string
		err error
	}{
		{
			name: "Barnaul",
			value: "Barnaul",
			err: nil,
		},
		{
			name: "ru_Barnaul",
			value: "Барнаул",
			err: SymbolsError,
		},
		{
			name: "Random symbols",
			value: "!!!///",
			err: SymbolsError,
		},
	}
	for _, tc := range cases  {
		t.Run(tc.name, func(t *testing.T) {
			err := InputCheck(tc.value)
			if err != tc.err {
				t.Errorf("%q failed, get: %q", tc.name, err)
			}
		})
	}
}