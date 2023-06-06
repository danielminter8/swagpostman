package swagpostman

type mockPostman struct{}

func NewMockPostman() *mockPostman {
	return &mockPostman{}
}

func (c *mockPostman) UpdateCollection(collectionID, swaggerString string) error {
	// Implement the mock behavior for updating the Postman collection
	return nil
}

type mockSwagger struct{}

func (f *mockSwagger) Get(url string) (string, error) {
	// Implement the mock behavior for fetching Swagger data
	return "mockSwaggerString", nil
}

func NewMockSwagger() *mockSwagger {
	return &mockSwagger{}
}
