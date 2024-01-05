package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	log.Println("Starting Application...")
	router := gin.Default()

	log.Println("Creating Route...")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {"status" : "Application Running." })
	})
	log.Println("Created Route...")
	log.Println("Application Started...")
	router.Run(":8085")
}