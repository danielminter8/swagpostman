package swagpostman

import "testing"

func TestPostmanSwagSync(t *testing.T) {
	apiKey := "your-api-key"

	syncer := NewSyncer(apiKey)

	syncer.Postman = NewMockPostman()
	syncer.Swagger = NewMockSwagger()

	collectionID := "your-collection-id"
	swaggerURL := "your-swagger-url"

	err := syncer.Sync(collectionID, swaggerURL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

}
