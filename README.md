# hexago-api-go
Welcome to the documentation for **hexago-api-go**. This API is built using the Go programming language and provides various endpoints for interacting with the application.

## Getting Started

### Prerequisites

Before you begin, make sure you have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

### Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/giocnidev/hexago-api-go.git
    cd hexago-api-go

2. Run `go mod tidy` to ensure that all dependencies are correctly managed:
    ```shell
    go mod tidy

3. BuilRund the project:
    ```shell
    go run cmd/api/main.go 

## API Endpoints

##### Get product

Retrieve a specific product.

- **URL**: `/api/v1/products/{productId}`
- **Method**: `GET`
- **URL Params**:
  - Required:
    - `productId=[string]`: The ID of the product to search
- **Request Body**:
  ```json
  {
    "id": "product_id",
    "name": "product name",
    "description": "product description",
    "price": 100
  }


### Create product

Create a product.

- **URL**: `/api/v1/products`
- **Method**: `POST`
- **URL Params**: none
- **Request Body**:
  ```json
  {
    "id": "product_id",
    "name": "product name",
    "description": "product description",
    "price": 100
  }



