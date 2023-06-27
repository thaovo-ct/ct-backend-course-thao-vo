package ad_listing

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"encoding/json"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)
type Option = func(c *client)

func NewClient(opts ...Option) *client {
	// TODO #4 refactor NewClient using functional options
	c := &client{
		httpClient: http.DefaultClient,
		baseUrl:    BaseUrl,
		retryTimes: 0,
		logger:     log.Default(),
	}
	
	for _, o := range opts {
		o(c)
	}

	return c
}

func WithBaseUrl(baseUrl string) Option {
	return func(c *client) {
		c.baseUrl = baseUrl
	}
}

func WithRetryTimes(retryTimes int) Option {
	return func(c *client) {
		c.retryTimes = retryTimes
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(c *client) {
		c.logger = logger
	}
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}


type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}

func (c *client) GetAdByCate(ctx context.Context, cate string) (*AdsResponse, error) {
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request - Cate %v, Duration: %v", cate, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, cate)

	// TODO #3 implement retry if StatusCode = 5xx
	var resp *http.Response
	var err error

	retryTime := c.retryTimes

	for retryTime >= 0{
		resp, err = c.httpClient.Get(url)
		if err != nil {
			return nil, err
		} else if resp.StatusCode / 100 == 5 {
			retryTime--
			continue
		}
		break
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nResponse %v", string(b))

	var adResp AdsResponse
	// TODO #2 unmarshal json
	err = json.Unmarshal(b, &adResp)
	if err != nil {
		return nil, err
	}
	return &adResp, nil
}

type AdsResponse struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AdId int `json:"ad_id"`
	list_id int `json:"list_id"`
	account_name string `json:"account_name"`
	subject string `json:"subject"`
	list_time int `json:"list_time"`
	//TODO #1 Define struct
	// list_id , account_name, subject, list_time
}
