package async

import (
	"net/http"
	"time"
)

type Response struct {
	http *http.Response
	err  error
}

func Get(urls ...string) <-chan *Response {
	channel := make(chan *Response, len(urls))
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	for _, url := range urls {
		go func(url string) {
			response, err := client.Get(url)

			if err != nil {
				defer response.Body.Close()
			}

			channel <- &Response{response, err}
		}(url)
	}

	return channel
}
