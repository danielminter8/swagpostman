package swagpostman

import (
	"github.com/danielminter8/swagpostman/internal/postman"
	"github.com/danielminter8/swagpostman/internal/swagger"
)

type postmanSwag struct {
	Postman postman.Actions
	Swagger swagger.Actions
	ApiKey  string
}

func NewSyncer(apiKey string) *postmanSwag {
	return &postmanSwag{
		ApiKey: apiKey,
	}
}

func (sg *postmanSwag) Sync(collectionID string, swaggerURL string) error {
	swaggerString, err := sg.Swagger.Get(swaggerURL)
	if err != nil {
		return err
	}

	err = sg.Postman.UpdateCollection(collectionID, swaggerString)
	if err != nil {
		return err
	}

	return nil
}
