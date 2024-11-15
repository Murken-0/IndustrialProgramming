package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id                string
	ManufacturerId    int
	ProductCategoryId int
	Name              string
	Price             float64
	Quantity          int
	ArticleNumber     int
	Description       string
	ImageUrl          string
}

var products = []Product{
	{Id: "1", ManufacturerId: 1, ProductCategoryId: 1, Name: "Samsung Galaxy S24", Price: 78000, Quantity: 10, ArticleNumber: 8100, Description: "Samsung Galaxy S24 — компактный смартфон из флагманской линейки, получивший мощный фирменный процессор, большой объем оперативной памяти, узнаваемый дизайн и продуманную эргономику. ", ImageUrl: "https://ir.ozone.ru/s3/multimedia-m/c1000/6900636406.jpg"},
	{Id: "2", ManufacturerId: 1, ProductCategoryId: 1, Name: "Samsung Galaxy S23 FE", Price: 36000, Quantity: 50, ArticleNumber: 7100, Description: "Смартфон Samsung Galaxy S23 FE - это новейшее устройство от известного южнокорейского производителя Samsung. Он предлагает пользователям высокую производительность и отличные функциональные возможности.", ImageUrl: "https://img.championat.com/i/d/o/168794395937683540.jpg"},
	{Id: "3", ManufacturerId: 2, ProductCategoryId: 2, Name: "Xiaomi Pad 6", Price: 26500, Quantity: 12, ArticleNumber: 5300, Description: "Xiaomi Pad 6: Стиль и производительность в одном устройстве Стильный дизайн, который впечатляет: Тонкий металлический корпус (6.51 мм) с элегантной фактурой", ImageUrl: "https://cdn1.ozone.ru/s3/multimedia-c/6718970892.jpg"},
	{Id: "4", ManufacturerId: 2, ProductCategoryId: 2, Name: "Xiaomi Pad 6 Pro", Price: 49000, Quantity: 5, ArticleNumber: 1000, Description: "Планшет Xiaomi Pad 6S Pro 8/256 GB - это устройство, которое сочетает в себе высокую производительность и стильный дизайн.", ImageUrl: "https://www.gizmochina.com/wp-content/uploads/2023/06/Xiaomi-Pad-6-Featured-scaled.jpg"},
	{Id: "5", ManufacturerId: 3, ProductCategoryId: 3, Name: "Asus TUF Gaming A15", Price: 85000, Quantity: 10, ArticleNumber: 1100, Description: "Игровой ноутбук ASUS TUF Gaming A15 FA507NU-LP141 обеспечивает впечатляющую производительность для игр и других задач.", ImageUrl: "https://avatars.mds.yandex.net/i?id=a03c8b88bfe7858a4bfea3daa18d8fcf_l-6998621-images-thumbs&n=13"},
	{Id: "6", ManufacturerId: 3, ProductCategoryId: 3, Name: "Asus ROG Zephyrus G15", Price: 15000, Quantity: 5, ArticleNumber: 19000, Description: "Мощный и портативный, ноутбук ROG Zephyrus G15 представляет собой игровую платформу на базе операционной системы Windows 10, выполненную в ультратонком корпусе весом всего 1,9 кг.", ImageUrl: "https://static.onlinetrade.ru/img/items/m/noutbuk_asus_rog_zephyrus_m15_gu502lv_az105t_15.6_fhd_ag_ips_240hz_i7_10750h_16gb_1024gb_ssd_nodvd_rtx_2060_6gb_w10_black_1459559_3.jpg"},
	{Id: "7", ManufacturerId: 3, ProductCategoryId: 3, Name: "Asus Vivobook 17", Price: 41500, Quantity: 12, ArticleNumber: 900, Description: "ASUS Vivobook 17 X1704ZA-AU341 90NB10F2-M00DD0 — производительный ноутбук в ударопрочном корпусе с активной системой охлаждения.", ImageUrl: "https://avatars.mds.yandex.net/i?id=b40b973a43129287ac2cfdb6ff688283_l-6458590-images-thumbs&n=13"},
	{Id: "8", ManufacturerId: 4, ProductCategoryId: 4, Name: "Sony WH-1000XM4", Price: 25000, Quantity: 25, ArticleNumber: 3400, Description: "Беспроводные наушники Sony WH-1000XM4 — это флагманские наушники с активным шумоподавлением и высоким качеством звука.", ImageUrl: "https://ir.ozone.ru/s3/multimedia-2/c1000/6766801130.jpg"},
	{Id: "9", ManufacturerId: 4, ProductCategoryId: 4, Name: "Sony WH-CH720N", Price: 8000, Quantity: 50, ArticleNumber: 9000, Description: "Благодаря технологии шумоподавления, легкой конструкции и длительному времени работы от аккумулятора вы сможете наслаждаться музыкой дольше и без отвлекающих окружающих звуков.", ImageUrl: "https://cdn1.ozone.ru/s3/multimedia-x/c600/6670694193.jpg"},
	{Id: "10", ManufacturerId: 5, ProductCategoryId: 3, Name: "Apple MacBook Air 13 Retina", Price: 65000, Quantity: 10, ArticleNumber: 7900, Description: "Самый тонкий и лёгкий ноутбук Apple MacBook Air 13 model: A2337 теперь стал суперсильным благодаря чипу Apple M1.", ImageUrl: "https://msk.aura-rent.ru/wp-content/uploads/2020/02/noutbook-air-2.2.jpg"},
	{Id: "11", ManufacturerId: 5, ProductCategoryId: 3, Name: "Apple MacBook Pro 14 2023", Price: 135000, Quantity: 7, ArticleNumber: 14100, Description: "Ноутбук Apple Macbook Pro 14 M3 8/512 Silver (MR7J3) – мощный и стильный помощник для работы и развлечений.", ImageUrl: "https://mtscdn.ru/upload/iblock/ee8/mbp14_silver_gallery7_202301.jpg"},
	{Id: "12", ManufacturerId: 5, ProductCategoryId: 5, Name: "Apple iMac 24", Price: 190000, Quantity: 10, ArticleNumber: 410, Description: "Моноблок Apple iMac 24 2023 года - это мощный и стильный компьютер, который станет незаменимым помощником в вашей повседневной работе.", ImageUrl: "https://avatars.mds.yandex.net/get-mpic/3732535/2a0000018a6d36516ed5f119259f018168b1/orig"},
	{Id: "13", ManufacturerId: 6, ProductCategoryId: 6, Name: "Nintendo Switch OLED", Price: 25500, Quantity: 20, ArticleNumber: 16700, Description: "Консоль Nintendo Switch OLED с красочным 7-дюймовым экраном. При практически одинаковых размерах с Nintendo Switch консоль Nintendo Switch OLED отличается более крупным 7-дюймовым OLED-экраном с глубокими цветами и высоким контрастом.", ImageUrl: "https://cdn1.ozone.ru/s3/multimedia-s/6080757748.jpg"},
	{Id: "14", ManufacturerId: 7, ProductCategoryId: 6, Name: "Sony PlayStation 5 Slim", Price: 47000, Quantity: 16, ArticleNumber: 4500, Description: "Игровая приставка Sony PlayStation 5 Slim: улучшенный дизайн и расширенные возможности хранения данных Обновленная версия популярной консоли PlayStation 5", ImageUrl: "https://appmistore.ru/upload/iblock/c4e/25bnswfucwuectbrw1s9vlgvr4bkswud.webp"},
	{Id: "15", ManufacturerId: 8, ProductCategoryId: 6, Name: "Microsoft Xbox Series S", Price: 34000, Quantity: 25, ArticleNumber: 19300, Description: "Игровая консоль Microsoft Xbox Series S рассчитана на использование игр, загружаемых из цифровой библиотеки. ", ImageUrl: "https://media.wired.co.uk/photos/606d9dbb20fc96acca6d3a5a/1:1/w_2000,h_2000,c_limit/3.jpg"},
	{Id: "16", ManufacturerId: 9, ProductCategoryId: 7, Name: "JBL Charge 5", Price: 13000, Quantity: 40, ArticleNumber: 56000, Description: "Charge 5 - портативная колонка от JBL, предназначенная для использования в любых условиях. Она имеет защиту от воды и пыли по стандарту IP67, что позволяет использовать ее на пляже или в походе.", ImageUrl: "https://kazandigital.ru/uploaded/images/abouts/157610-speakers-review-jbl-charge-5-review-image1-4bvjkgsxy5.jpg"},
	{Id: "17", ManufacturerId: 9, ProductCategoryId: 7, Name: "JBL Flip 6", Price: 9000, Quantity: 40, ArticleNumber: 109000, Description: "Страна-производитель Китай Общие параметры Тип портативная колонка Модель JBL Flip 6 Код производителя [JBLFLIP6BLK] Основной цвет черный Акустические характеристики", ImageUrl: "https://avatars.mds.yandex.net/get-mpic/2017233/2a0000018de355c9d630bec9a88e2b702591/optimize"},
	{Id: "18", ManufacturerId: 10, ProductCategoryId: 8, Name: "NVIDIA GeForce RTX 3060 Dual LHR", Price: 29000, Quantity: 20, ArticleNumber: 67100, Description: "Видеокарта Palit GeForce RTX 3060 Dual 12G поможет тебе получить стабильный FPS выше 60 кадров в секунду при максимальных настройках графики и разрешении Full HD.", ImageUrl: "https://avatars.mds.yandex.net/get-mpic/10352132/2a0000018ec730cbd98af51c3c2abe51c47d/optimize"},
	{Id: "19", ManufacturerId: 11, ProductCategoryId: 8, Name: "AMD Radeon RX 6600 PULSE", Price: 36000, Quantity: 41, ArticleNumber: 41100, Description: "Техпроцесс: 7 нм; Тип видеокарты: игровая; Графический процессор: Radeon RX 6600;", ImageUrl: "https://avatars.mds.yandex.net/get-mpic/4120567/2a000001922a9d36a89c564a9f2f29901969/optimize"},
	{Id: "20", ManufacturerId: 1, ProductCategoryId: 9, Name: "Samsung Galaxy Watch 6", Price: 19000, Quantity: 40, ArticleNumber: 9900, Description: "Смарт-часы Samsung Galaxy Watch6 44 мм Silver (SM-R940) обладают дисплеем диагональю 1,47 дюйма и разрешением 480x480 пикселей — увеличить экран и добавить пространства для свайпов позволила тонкая рамка.", ImageUrl: "https://avatars.mds.yandex.net/get-mpic/10352132/2a0000018c5f0c1ddd3b729147a231ecc7b7/optimize"},
}

func main() {
	router := gin.Default()

	// Получение всех товаров
	router.GET("/products", getProducts)

	// Получение товара по ID
	router.GET("/products/:id", getProductByID)

	// Создание нового товара
	router.POST("/products", createProduct)

	// Обновление существующего товара
	router.PUT("/products/:id", updateProduct)

	// Удаление товара
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.Id == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, product := range products {
		if product.Id == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
