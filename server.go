package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "pen", Price: 75.20},
	{ID: "2", Name: "book", Price: 50.80},
	{ID: "3", Name: "notebook", Price: 90.00},
	{ID: "4", Name: "paper", Price: 25.60},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProduct)
	router.GET("/products/:id", getProductByID)
	router.GET("/products/search", getProductByQuery)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	server.ListenAndServe()
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
}

func getProductByQuery(c *gin.Context) {
	name := c.Query("name")
	var results []Product

	for _, product := range products {
		if product.Name == name {
			results = append(results, product)
			c.JSON(http.StatusOK, results)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
	
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Product deleted",
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Product not found",
	})

}
