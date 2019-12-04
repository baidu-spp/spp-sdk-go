package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   4 * time.Second,
				KeepAlive: 90 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 60,
			IdleConnTimeout:       time.Duration(90) * time.Second,
		},
		Timeout: 60 * time.Second}
)

func HttpPost(uri string, header, queryArgs map[string]string, body []byte) (*http.Response, []byte, error) {
	if len(queryArgs) > 0 {
		u, err := url.Parse(uri)
		if err != nil {
			return nil, nil, err
		}
		q := u.Query()
		for k, v := range queryArgs {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
		uri = u.String()
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("response[%s] status[%d] data[%s]", uri, resp.StatusCode, string(b))
		return resp, b, err
	}
	return resp, b, nil
}
