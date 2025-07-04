# Update Product Microservice

This microservice allows updating product details in the ToyShop platform. It is part of the product management domain and interacts with a MongoDB database.

## Technologies Used

- Go (Golang)
- Gin (web framework)
- MongoDB
- Docker
- GitHub Actions

## Getting Started

### Prerequisites

- Go >= 1.18
- MongoDB
- Git

### Installation

```bash
git clone https://github.com/andrespaida/update_product.git
cd update_product
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
PORT=4003
MONGO_URI=mongodb://your_mongo_host:27017
DB_NAME=toyshop_db
COLLECTION_NAME=products
```

### Running the Service

```bash
go run main.go
```

The service will be running at `http://localhost:4003`.

## Available Endpoint

### PUT `/products/:id`

Updates the product with the specified ID.

#### Request body (JSON):

```json
{
  "name": "Updated Product",
  "description": "Updated description",
  "price": 25.99,
  "stock": 30,
  "category": "New Category"
}
```

#### Example Response:

```json
{
  "message": "Product updated successfully"
}
```

## Docker

To build and run the service using Docker:

```bash
docker build -t update-product .
docker run -p 4003:4003 --env-file .env update-product
```

## GitHub Actions Deployment

This project includes a GitHub Actions workflow for automatic deployment to an EC2 instance. Configure the following secrets in your GitHub repository:

- `EC2_HOST`
- `EC2_USERNAME`
- `EC2_KEY`
- `EC2_PORT` (optional)

## License

This project is licensed under the MIT License.
