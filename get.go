package async

import (
	"net/http"
	"time"
)

type Response struct {
	http *http.Response
	err  error
}

func Get(url string) (*http.Response, error) {
	channel := make(chan Response, 1)

	go func() {
		client := &http.Client{
			Timeout: 60 * time.Second,
		}

		response, err := client.Get(url)
		response.Body.Close()

		channel <- Response{response, err}
	}()

	response := <-channel

	return response.http, response.err
}
