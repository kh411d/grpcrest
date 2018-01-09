package main

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	buf := bytes.NewBufferString(`
		{
			"name":"John Doe"
		}
	`)
	client.Post("http://localhost:9090", "application/json", buf)

}
