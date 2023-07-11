package ad_listing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)

// TODO #4 refactor NewClient using functional options
func NewClient(opt ...Option) *client {
    c := &client{
		httpClient: http.DefaultClient,
		baseUrl: BaseUrl,
		retryTimes: 0,
		logger: log.Default(),
	}
    for _, o := range opt {
        o(c)
    }
    return c
}

func WithBaseUrl(url string) Option {
    return func(c *client) {
        c.baseUrl = url
    }
}

func WithBaseClient(base *http.Client) Option {
    return func(c *client) {
        c.httpClient = base
    }
}

func WithRetryTimes(retryTime int) Option {
    return func(c *client) {
        c.retryTimes = retryTime
    }
}

func WithLogger(logger *log.Logger) Option {
    return func(c *client) {
        c.logger = logger
    }
}

type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}

type Option func (*client)

func (c *client) GetAdByCate(ctx context.Context, cate string) (*AdsResponse, error) {
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request - Cate %v, Duration: %v", cate, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, cate)
	// TODO #3 implement retry if StatusCode = 5xx
	var resp *http.Response
	var err error
	for (c.retryTimes > 0) {
		resp, err = c.httpClient.Get(url)
		if (err != nil) {
			return nil, err
		} else if resp.StatusCode >= 500 {
			c.retryTimes--
			continue
		}
		break
	}
	
	if err != nil {
		return nil, err
	}
	
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nResponse %v", string(b))

	var adResp AdsResponse
	// TODO #2 unmarshal json
	err = json.Unmarshal(b, &adResp)

	if (err != nil){
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
	//TODO #1 Define struct
	ListId int `json:"list_id"`
	AccountName string `json:"account_name"`
	Subject string `json:"subject"`
	ListTime int `json:"list_time"`
	// list_id , account_name, subject, list_time
}