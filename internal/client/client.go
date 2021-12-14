package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	m "translateapp/internal/models"
)

const BaseURLLibre = "http://libretranslate:5000"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:    BaseURLLibre,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) GetLanguages(ctx context.Context) (*m.ClientSuccessResponse, *m.ClientErrResponse) {
	var errResp m.ClientErrResponse
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/languages", c.BaseURL), nil)
	if err != nil {
		errResp.Message = fmt.Sprintf("%s", err)
		errResp.Code = 500
		return nil, &errResp
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		errResp.Message = fmt.Sprintf("%s", err)
		errResp.Code = res.StatusCode
		return nil, &errResp
	}
	defer res.Body.Close()
	bytes, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		errResp.Message = fmt.Sprintf("%s", errRead)
		errResp.Code = 500
		return nil, &errResp
	}
	var response m.ClientSuccessResponse
	response.Code = res.StatusCode
	response.Data = string(bytes)
	return &response, nil
}
