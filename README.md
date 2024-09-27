# Simple API with Gin

This project is an introduction to the concepts of a web API with Go and Gin Web Framework.


## 📋 Description

The API allows you to create, list, search, update, and delete products. This is a basic project to learn how to set up routes, handle requests and responses, and work with JSON in a Go application.

## 🚀 Getting Started

### 💻 Prerequisites
- [Golang](https://go.dev/doc/install) installed on your machine

### 🔧 How to Build and Run the Project

**To build the project**, use the command:

```go
go build server.go
```

**run the project**
**To run the server**, execute the generated file:

```go
./server
```

By default, the server will start on port 8000.

# 📪 Endpoints

### GET /products
List all registered products.

- **Request:** `localhost:8000/products`

### GET /products/{id}
Get details of a specific product by ID.

- **Request:** `localhost:8000/products/1`


### GET /products/search
Search for a product by name.

- **Query parameter:** product name
- **Request:** `localhost:8000/products/search?name=book`


### POST /products
Create a new product.
- **Request:** `localhost:8000/products` 
- **Request body:**

```json
{
  "id": 5,
  "name": "table",
  "price": 125.60
}
```

### PUT /products/{id}

Update the details of an existing product by ID.

- **Request:** `localhost:8000/products/1`
- **Request body:**

```json
{
   "id": 1,
  "name": "pen",
  "price": 90.45
}
```

### DELETE /products/{id}
Delete a specific product by ID.

- **Request:** `localhost:8000/products/2`
