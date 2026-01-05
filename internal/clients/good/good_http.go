package orderhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"telegram/internal/model/good"
	"time"
)

type GoodHTTPClient struct {
	baseURL string
	client  *http.Client
}

func InitgoodHTTPClient(baseURL string) *GoodHTTPClient {
	return &GoodHTTPClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (o *GoodHTTPClient) AddGood(ctx context.Context, Id string) (good.Good, error) {
	var good good.Good

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return good, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return good, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return good, fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return gooder, err
	}

	return gooder, nil
}

func (o *GoodHTTPClient) DeleteGood(ctx context.Context, Id string) error {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return err
	}

	return nil
}

func (o *GoodHTTPClient) GetAllGoods(ctx context.Context) ([]good.Good, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL), nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder []good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return gooder, err
	}

	return gooder, nil
}

func (o *GoodHTTPClient) GetByIdOrder(ctx context.Context, Id string) (good.Good, error) {
	var good good.Good

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return good, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return good, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return good, fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return gooder, err
	}

	return gooder, nil
}

func (o *GoodHTTPClient) UpdateOrder(ctx context.Context, Id string) error {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/good?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return err
	}

	return nil
}

func (o *GoodHTTPClient) IsAvaliableForOrder(ctx context.Context, Id string) (bool, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid", o.baseURL), nil,
	)
	if err != nil {
		return false, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return false, err
	}

	return true, nil
}

func (o *GoodHTTPClient) RestOfGood(ctx context.Context, Id string) (int64, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/good?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return 0, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("good service returned %d", response.StatusCode)
	}

	var gooder good.Good
	if err := json.NewDecoder(response.Body).Decode(&gooder); err != nil {
		return 0, err
	}

	return good.Good.Rest, nil
}
