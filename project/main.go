package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Manufacturer struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Category struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Product struct {
	ID           string       `json:"id"`
	Code         string       `json:"code"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Price        string       `json:"price"`
	Image        string       `json:"image"`
	Manufacturer Manufacturer `json:"manufacturer"`
	Category     Category     `json:"category"`
}

var products = []Product{
	{
		ID:          "1",
		Code:        "ABC345",
		Title:       "A",
		Description: "B",
		Price:       "155.6",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-w/wc300/7000867676.jpg",
		Manufacturer: Manufacturer{
			ID:          "1",
			Title:       "A",
			Description: "ssasa",
		},
		Category: Category{
			ID:    "1",
			Code:  "EL",
			Title: "A",
		},
	},
	{
		ID:          "2",
		Code:        "ABC346",
		Title:       "C",
		Description: "D",
		Price:       "142345.6",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-w/wc300/7000867676.jpg",
		Manufacturer: Manufacturer{
			ID:          "1",
			Title:       "A",
			Description: "ssasa",
		},
		Category: Category{
			ID:    "1",
			Code:  "CORM",
			Title: "F",
		},
	},
}

func main() {
	router := gin.Default()

	// Получение всех книг
	router.GET("/books", getProducts)

	// Получение книги по ID
	router.GET("/books/:id", getProductByID)

	// Создание новой книги
	router.POST("/books", createBook)

	// Обновление существующей книги
	router.PUT("/books/:id", updateBook)

	// Удаление книги
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
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

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
