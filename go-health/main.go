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
	
	router.Run(":8085")
	log.Println("Application Started...")
}