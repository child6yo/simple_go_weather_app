package main

import "testing"

func TestParceAnyResponce(t *testing.T) {
	cases := []struct {
		name string
		inter Responces
		value []byte
		err error
	}{
		{
			name: "empty",
			inter: &CoordinatesAPIResponce{},
			value: []byte{},
			err: ParceError,
		},
	}
	for _, tc := range cases  {
		t.Run(tc.name, func(t *testing.T) {
			err := ParceAnyResponce(tc.inter, tc.value)
			if err != tc.err {
				t.Errorf("%q failed", tc.name)
			}
		})
	}
}