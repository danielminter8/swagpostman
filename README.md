# SwagPostman - Postman Swagger Sync

[![Go Reference](https://pkg.go.dev/badge/github.com/danielminter8/swagpostman.svg)](https://pkg.go.dev/github.com/danielminter8/swagpostman)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielminter8/swagpostman)](https://goreportcard.com/report/github.com/danielminter8/swagpostman)
[![License](https://img.shields.io/github/license/danielminter8/swagpostman)](https://github.com/danielminter8/swagpostman/blob/main/LICENSE)

SwagPostman is a Go package that simplifies the process of synchronizing Swagger documentation with a Postman collection. It provides an easy-to-use interface for developers to update their collections with the latest Swagger specifications, ensuring accurate and up-to-date API documentation in their Postman workflows.

Maintaining consistency between Swagger documentation and Postman collections can be a time-consuming and error-prone task. With Postman Swagger Sync, you can automate this process, saving valuable time and reducing the chances of discrepancies between your API specifications and Postman environments.

Key Features:

- Sync Swagger Documentation: The package enables you to effortlessly synchronize your Swagger documentation with a Postman collection. Simply provide your Postman API key, the collection ID, and the path to your Swagger file, and let the package handle the rest.
- Efficient and Reliable: Postman Swagger Sync leverages Go's standard library and industry-standard practices, making it a dependable choice for syncing your Swagger documentation.
- Intuitive API: The package offers a straightforward API that is easy to understand and use. With clear method names and intuitive parameters, developers can quickly integrate the package into their projects and streamline their Swagger-to-Postman synchronization process.
- Error Handling: Postman Swagger Sync incorporates comprehensive error handling to provide informative feedback and handle unexpected scenarios. It returns meaningful error messages, enabling developers to troubleshoot and resolve any issues that may arise during the synchronization process.

Whether you are building a new API or maintaining an existing one, keeping your Swagger documentation and Postman collections in sync is crucial. With Swagpostman, you can automate this synchronization and ensure that your Postman environments always reflect the latest API specifications.

## Installation

```bash
go get github.com/danielminter8/swagpostman
```

## Usage

```go
import (
	"log"

	"github.com/danielminter8/swagpostman"
)

func main() {
	// Set your Postman API key
	apiKey := "YOUR_POSTMAN_API_KEY"

	// Set your collection ID
	collectionID := "YOUR_POSTMAN_COLLECTION_ID"

    // Set your swagger source
    swaggerURL := "http://example.com/swagger/doc.json"

	// Create a new instance of the Syncer
	swagpostman := swagpostman.NewSyncer(apiKey)

	// Sync the Swagger documentation with the Postman collection
	err := swagpostman.Sync(collectionID, swaggerURL)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Swagger documentation synced with Postman collection successfully!")
}
```

## TODO

- [ ] Unit tests

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to create an issue or submit a pull request.

## License
This project is licensed under the [MIT License](https://spdx.org/licenses/MIT.html). <br>