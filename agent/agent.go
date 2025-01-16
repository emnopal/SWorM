package agent

import "github.com/go-resty/resty/v2"

func Send(url string, method string) {
	client := resty.New()
	request := client.R()

	if method == "GET" {
		request.Get(url)
	} else if method == "POST" {
		request.Post(url)
	}
	// WIP
}
