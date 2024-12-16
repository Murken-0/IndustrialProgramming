package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID             int     `gorm:"primaryKey;column:id" json:"id"`
	Code           string  `gorm:"column:code" json:"code"`
	Title          string  `gorm:"column:title" json:"title"`
	Description    string  `gorm:"column:description" json:"description"`
	Price          float64 `gorm:"column:price" json:"price"`
	Preview        string  `gorm:"column:preview" json:"preview"`
	ManufacturerId int     `gorm:"column:manufacturer_id" json:"manufacturer_id"`
	CategoryId     int     `gorm:"column:category_id" json:"category_id"`
}

type CartProducts struct {
	ProductID int `gorm:"primaryKey;column:product_id" json:"product_id"`
	UserID    int `gorm:"primaryKey;column:user_id" json:"user_id"`
	Quantity  int `gorm:"column:quantity" json:"quantity"`
}

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userId"`
	jwt.StandardClaims
}

type User struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	Login       string `gorm:"column:login;" json:"login"`
	Password    string `gorm:"column:password;" json:"password"`
	Email       string `gorm:"column:email;" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;" json:"phoneNumber"`
	IsActive    bool   `gorm:"column:is_active" json:"is_active"`
}

var router = gin.Default()

var db *gorm.DB

func main() {
	initDB()
	router.POST("/login", login)
	router.POST("/refresh", refreshToken)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.POST("/products", createProduct)
		protected.PUT("/products/:id", updateProduct)
		protected.DELETE("/products/:id", deleteProduct)
		protected.GET("/cart", getCart)
		protected.POST("/cart", addToCart)
		protected.DELETE("/cart/:productId", deleteFromCart)
	}

	router.Run(":8080")
}

func initDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=industrial port=5432 sslmode=disable search_path=public"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connected successfully")
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CartProducts{})
	log.Println("Database initialized")
}

func generateToken(username string, userID int) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var user User
	if err := db.Where("login = ? AND password = ?", creds.Login, creds.Password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		}
		return
	}

	token, err := generateToken(creds.Login, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}

func refreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				newToken, err := generateToken(claims.Username, claims.UserID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "could not refresh token"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"token": newToken})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "token is still valid"})
}

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func getCart(c *gin.Context) {
	userID := c.GetInt("userId")
	var cartItems []CartProducts
	if err := db.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func addToCart(c *gin.Context) {
	var newItem CartProducts

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var existingItem CartProducts
	result := db.Where("product_id = ? AND user_id = ?", newItem.ProductID, newItem.UserID).First(&existingItem)

	if result.RowsAffected > 0 {
		existingItem.Quantity += newItem.Quantity
		if err := db.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update cart item"})
			return
		}
		c.JSON(http.StatusOK, existingItem)
		return
	}

	if err := db.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to add item to cart"})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func deleteFromCart(c *gin.Context) {
	productID := c.Param("productId")
	userID := c.GetInt("userId")

	var item CartProducts
	if err := db.Where("product_id = ? AND user_id = ?", productID, userID).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "product not found in cart"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch cart item"})
		}
		return
	}

	if err := db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to remove product from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product removed from cart"})
}
