package orderhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"telegram/internal/model/order"
	"time"
)

type OrderHTTPClient struct {
	baseURL string
	client  *http.Client
}

func InitOrderHTTPClient(baseURL string) *OrderHTTPClient {
	return &OrderHTTPClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (o *OrderHTTPClient) AddOrder(ctx context.Context, Id string) (order.Order, error) {
	var ord order.Order

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return ord, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return ord, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ord, fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return order, err
	}

	return order, nil
}

func (o *OrderHTTPClient) DeleteOrder(ctx context.Context, Id string) error {
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
		return fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return err
	}

	return nil
}

func (o *OrderHTTPClient) GetAllOrders(ctx context.Context) ([]order.Order, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid", o.baseURL), nil,
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
		return nil, fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order []order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return order, err
	}

	return order, nil
}

func (o *OrderHTTPClient) GetByIdOrder(ctx context.Context, Id string) (order.Order, error) {
	var ord order.Order

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return ord, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return ord, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ord, fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return order, err
	}

	return order, nil
}

func (o *OrderHTTPClient) UpdateOrder(ctx context.Context, Id string) error {
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
		return fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return err
	}

	return nil
}

func (o *OrderHTTPClient) IsOrderCreated(ctx context.Context, Id string) (bool, error) {
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
		return false, fmt.Errorf("order service returned %d", response.StatusCode)
	}

	var order order.Order
	if err := json.NewDecoder(response.Body).Decode(&order); err != nil {
		return false, err
	}

	return true, nil
}
