package main

import "testing"

func TestParceAnyResponse(t *testing.T) {
	cases := []struct {
		name string
		inter Responses
		value []byte
		err error
	}{
		{
			name: "empty",
			inter: &CoordinatesAPIResponse{},
			value: []byte{},
			err: ParseError,
		},
	}
	for _, tc := range cases  {
		t.Run(tc.name, func(t *testing.T) {
			err := ParceAnyResponse(tc.inter, tc.value)
			if err != tc.err {
				t.Errorf("%q failed", tc.name)
			}
		})
	}
}