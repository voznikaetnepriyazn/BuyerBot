package customerhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"telegram/internal/model/customer"
	"time"
)

type CustomerHTTPClient struct {
	baseURL string
	client  *http.Client
}

func InitCustomerHTTPClient(baseURL string) *CustomerHTTPClient {
	return &CustomerHTTPClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (o *CustomerHTTPClient) AddCustomer(ctx context.Context, Id string) (customer.Customer, error) {
	var cus customer.Customer

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return cus, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return cus, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return cus, fmt.Errorf("customer service returned %d", response.StatusCode)
	}

	var customer customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return customer, err
	}

	return customer, nil
}

func (o *CustomerHTTPClient) DeleteCustomer(ctx context.Context, Id string) error {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/customer?userid=%d", o.baseURL, Id), nil,
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
		return fmt.Errorf("customer service returned %d", response.StatusCode)
	}

	var customer customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return err
	}

	return nil
}

func (o *CustomerHTTPClient) GetAllCustomers(ctx context.Context) ([]customer.Customer, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/customer?userid", o.baseURL), nil,
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

	var customer []customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return customer, err
	}

	return customer, nil
}

func (o *CustomerHTTPClient) GetByIdCustomer(ctx context.Context, Id string) (customer.Customer, error) {
	var cus customer.Customer

	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/order?userid=%d", o.baseURL, Id), nil,
	)
	if err != nil {
		return cus, err
	}

	response, err := o.client.Do(request)
	if err != nil {
		return cus, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return cus, fmt.Errorf("customer service returned %d", response.StatusCode)
	}

	var customer customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return customer, err
	}

	return customer, nil
}

func (o *CustomerHTTPClient) UpdateCustomer(ctx context.Context, Id string) error {
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
		return fmt.Errorf("customer service returned %d", response.StatusCode)
	}

	var customer customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return err
	}

	return nil
}

func (o *CustomerHTTPClient) IsCustomerCreated(ctx context.Context, Id string) (bool, error) {
	request, err := http.NewRequestWithContext(
		ctx, "GET", fmt.Sprintf("%s/api/customer?userid=%d", o.baseURL, Id), nil,
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
		return false, fmt.Errorf("customer service returned %d", response.StatusCode)
	}

	var customer customer.Customer
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return false, err
	}

	return true, nil
}
