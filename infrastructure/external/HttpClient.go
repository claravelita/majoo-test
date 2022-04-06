package external

import (
	"net"
	"net/http"
	"time"
)

type httpClient struct {
	Client http.Client
}

func setTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
}

func setHttpClient() http.Client {
	return http.Client{
		Timeout:   time.Second * 10,
		Transport: setTransport(),
	}
}

func GetHttpClient() httpClient {
	return httpClient{
		setHttpClient(),
	}
}
