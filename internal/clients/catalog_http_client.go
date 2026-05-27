package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ordersModel "example.com/m/v4/internal/models/orders"
)

type CatalogHTTPClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCatalogHTTPClient(baseURL string) *CatalogHTTPClient {
	return &CatalogHTTPClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *CatalogHTTPClient) GetProductByID(
	ctx context.Context,
	productID uint,
) (*ordersModel.ProductSnapshot, error) {
	url := fmt.Sprintf("%s/api/product/%d", c.baseURL, productID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("product %d not found", productID)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}

	var product ordersModel.ProductSnapshot
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, fmt.Errorf("invalid product response from catalog service")
	}

	return &product, nil
}
