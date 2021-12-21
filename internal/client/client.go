package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"translateapp/internal/models"
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

func (c *Client) GetLanguages(ctx context.Context) (*models.LanguageList, *models.ClientErrResponse) {
	var errResp models.ClientErrResponse
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/languages", c.BaseURL), nil)
	if err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		errResp.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&errResp.Message); err != nil {
			errResp.Message.Message = err.Error()
			return nil, &errResp
		}
		return nil, &errResp
	}
	var langList models.LanguageList
	if err := json.NewDecoder(res.Body).Decode(&langList.Languages); err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	return &langList, nil
}

func (c *Client) Translate(ctx context.Context, input models.Input) (*models.Translation, *models.ClientErrResponse) {
	var errResp models.ClientErrResponse
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(input)
	if err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/translate", c.BaseURL), &body)
	if err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	defer res.Body.Close()
	if res.StatusCode != 200 { //nolint:wsl
		errResp.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&errResp.Message); err != nil {
			errResp.Message.Message = err.Error()
			return nil, &errResp
		}
		return nil, &errResp
	}
	var translation models.Translation
	//var translation models.Translation
	if err := json.NewDecoder(res.Body).Decode(&translation); err != nil {
		errResp.Message.Message = err.Error()
		errResp.Code = 500
		return nil, &errResp
	}
	return &translation, nil
}
