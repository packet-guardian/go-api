package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PGApi struct {
	url                *url.URL
	username, password string
}

func Connect(host string) (*PGApi, error) {
	hostUrl, err := url.Parse(host)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	pg := &PGApi{
		url: hostUrl,
	}
	return pg, nil
}

func (pg *PGApi) Login(username, password string) {
	pg.username = username
	pg.password = password
}

func (pg *PGApi) makeRequest(method, url string, params ...interface{}) (*http.Response, error) {
	url = "%s" + url
	params = append([]interface{}{pg.url.String()}, params...)
	fullURL := fmt.Sprintf(url, params...)

	client := &http.Client{}
	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(pg.username, pg.password)

	return client.Do(req)
}

type genericResponse struct {
	Message string
}

func (pg *PGApi) decodeGenericResponse(r io.Reader) *genericResponse {
	var resp genericResponse
	decoder := json.NewDecoder(r)
	decoder.Decode(&resp)
	return &resp
}
