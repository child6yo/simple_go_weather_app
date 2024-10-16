package main

import (
	"io"
	"net/http"
)

type Request struct {
	Url string
}

func (r *Request)GetResponse() ([]byte, error) {
	res, err := http.Get(r.Url)
	if err != nil {
		return []byte{}, RequestError
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, ReadError
	}
	return body, nil
}