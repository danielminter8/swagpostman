package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	baseURL          = "https://api.getpostman.com"
	collectionSchema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
)

var (
	errSwagPostmanSyncFailed = errors.New("error occured while trying to sync swagger docs with specified postman collection")
	errPostmanUnauthorized   = errors.New("postman API key not valid")
)

type Actions interface {
	UpdateCollection(collectionID, swaggerString string) error
}

type Client struct {
	ApiKey string
}

func (client *Client) UpdateCollection(ctx context.Context, collectionID, swaggerString string) error {

	// Create the API request to update the collection with Swagger docs
	url := fmt.Sprintf("%s/collections/%s", baseURL, collectionID)
	reqBody := Collection{
		Collection: struct {
			Info CollectionInfo `json:"info"`
			Item []Item         `json:"item"`
		}{
			Info: CollectionInfo{
				Schema: collectionSchema,
			},
			Item: []Item{
				{
					Name:     "Swagger Docs",
					Protocol: "swagger",
					Request: Request{
						Method: "GET",
						Header: []interface{}{},
						URL:    "{{url}}",
						Body: map[string]interface{}{
							"raw": swaggerString, // Set the Swagger data in the request body
						},
						Auth:    nil,
						Proxy:   nil,
						TLS:     nil,
						Version: "",
					},
					Response: []interface{}{
						map[string]interface{}{
							"status": "200",
							"code":   "Success",
						},
					},
				},
			},
		},
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return errSwagPostmanSyncFailed
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", client.ApiKey)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 401 {
		return errPostmanUnauthorized
	}

	return nil
}
