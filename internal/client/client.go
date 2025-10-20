package client

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/vadimklimov/cpi-mcp-server/internal/appinfo"
	"github.com/vadimklimov/cpi-mcp-server/internal/config"
	"golang.org/x/oauth2/clientcredentials"
	"resty.dev/v3"
)

type Client struct {
	*resty.Client
}

var (
	instance *Client
	once     sync.Once
)

func GetInstance(ctx context.Context) *Client {
	once.Do(func() {
		timeout, _ := time.ParseDuration(strconv.Itoa(config.Timeout()) + "s")

		oauthConfig := &clientcredentials.Config{
			TokenURL:     config.TokenURL().String(),
			ClientID:     config.ClientID(),
			ClientSecret: config.ClientSecret(),
		}

		httpClient := oauthConfig.Client(ctx)

		client := resty.NewWithClient(httpClient).
			SetContext(ctx).
			SetTimeout(timeout).
			SetBaseURL(config.BaseURL().String()).
			SetHeaders(map[string]string{
				"Accept":     "application/json",
				"User-Agent": appinfo.ID(),
			})

		instance = &Client{client}
	})

	return instance
}

func Get(request *resty.Request) (*resty.Response, error) {
	response, err := request.
		SetMethod(resty.MethodGet).
		Send()
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("failed to complete API request: response status: %s", response.Status())
	}

	return response, nil
}
