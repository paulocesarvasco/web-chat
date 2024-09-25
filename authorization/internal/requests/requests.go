package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"

	// "net"
	"net/http"
	"time"
)

type Requester interface{}

type request struct {
	client  *http.Client
	request *http.Request
}

func NewPostRequest(ctx context.Context, url string, payload any) (*request, error) {
	var rawPayload []byte
	switch p := payload.(type) {
	case []byte:
		rawPayload = p
	default:
		rp, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}
		rawPayload = rp
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		url, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}
	return &request{
		client:  defaultHTTPClient(),
		request: req,
	}, nil
}

func (r *request) AddBasicAuth(user, pass string) {
	r.request.SetBasicAuth(user, pass)
}

func (r *request) Execute() (int, []byte) {
	res, err := r.client.Do(r.request)
	if err != nil {
		log.Print("failed to execute request: ", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	rawPayload, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("failed to decode response payload: ", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	return res.StatusCode, rawPayload
}

func defaultHTTPClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},

		// MaxIdleConns:        10,
		// IdleConnTimeout:     30 * time.Second,
		// DisableKeepAlives:   false,
		// DisableCompression:  false,
		// MaxIdleConnsPerHost: 5,

		// DialContext: (&net.Dialer{
		//	Timeout:   5 * time.Second,
		//	KeepAlive: 30 * time.Second,
		// }).DialContext,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}
