package libretranslate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Logger     *zap.SugaredLogger
}

type ClientInterface interface {
	GetLanguages(ctx context.Context) (*LanguageList, error)
	Translate(ctx context.Context, input Input) (*Translation, error)
}

func NewClient(logger *zap.SugaredLogger, baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Logger:     logger,
	}
}

func (c *Client) GetLanguages(ctx context.Context) (*LanguageList, error) {
	var errResp CustomError

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/languages", c.BaseURL), nil)
	if err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	defer res.Body.Close()
	c.Logger.Info("Contacted Libretranslate")
	if res.StatusCode != 200 {
		type Keeper struct {
			Error string `json:"error"`
		}
		var keeper Keeper
		errResp.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&keeper); err != nil {
			errResp.Message = err.Error()
			return nil, &errResp
		}
		errResp.Message = keeper.Error
		return nil, &errResp
	}

	var langList LanguageList
	if err := json.NewDecoder(res.Body).Decode(&langList.Languages); err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	return &langList, nil
}

func (c *Client) Translate(ctx context.Context, input Input) (*Translation, error) {
	var errResp CustomError
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/translate", c.BaseURL), &body)
	if err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}

	defer res.Body.Close()
	c.Logger.Info("Contacted Libretranslate")

	if res.StatusCode != 200 {
		type Keeper struct {
			Error string `json:"error"`
		}
		var keeper Keeper
		errResp.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&keeper); err != nil {
			errResp.Message = err.Error()
			c.Logger.Errorf("%s", errResp.Error())
			return nil, &errResp
		}
		errResp.Message = keeper.Error
		return nil, &errResp
	}
	var translation Translation
	//var translation models.Translation
	if err := json.NewDecoder(res.Body).Decode(&translation); err != nil {
		errResp.Code = 500
		errResp.Message = err.Error()
		c.Logger.Errorf("%s", errResp.Error())
		return nil, &errResp
	}
	return &translation, nil
}
