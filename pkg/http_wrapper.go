/*
Pookie — simple http client.
Copyright © 2025 Toichuev Ulukbek t.ulukbek01@gmail.com

Licensed under the MIT License.
*/

package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type StatusErr struct {
	Code   int
	Status string
}

func (e StatusErr) Error() string {
	return "invalid response status: " + e.Status
}

type HttpWrapperAbstract interface {
	HttpGet(uri string, headers map[string]string, params map[string]string, respBody interface{}) error
}

type HttpWrapper struct {
	client *http.Client
}

func NewHttpWrapper(timeout int) *HttpWrapper {
	return &HttpWrapper{
		client: &http.Client{Timeout: time.Duration(timeout) * time.Second},
	}
}

func (hw HttpWrapper) HttpGetRaw(uri string, headers map[string]string, params map[string]string) (*http.Response, []byte, error) {
	req, err := createRequest(uri, headers, params)
	if err != nil {
		return nil, nil, err
	}

	resp, body, err := doRequest(*hw.client, req)
	if err != nil {
		return resp, body, err
	}

	return resp, body, nil
}

func (hw HttpWrapper) HttpGet(uri string, headers map[string]string, params map[string]string, respBody interface{}) (*http.Response, error) {
	req, err := createRequest(uri, headers, params)
	if err != nil {
		return nil, err
	}

	resp, body, err := doRequest(*hw.client, req)
	if err != nil {
		return resp, err
	}

	return resp, convertToJSON(body, respBody)
}

func createRequest(uri string, headers map[string]string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	urlParams := url.Values{}
	for key, val := range params {
		urlParams.Set(key, val)
	}
	req.URL.RawQuery = urlParams.Encode()
	return req, nil
}

func doRequest(client http.Client, req *http.Request) (*http.Response, []byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return resp, nil, fmt.Errorf("error performing request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, nil, StatusErr{resp.StatusCode, resp.Status}
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("error reading response body: %w", err)
	}
	return resp, body, nil
}

func convertToJSON(body []byte, respBody interface{}) error {
	err := json.Unmarshal(body, &respBody)
	if err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}
	return nil
}
