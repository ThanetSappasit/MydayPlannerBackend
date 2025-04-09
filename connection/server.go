package connection

import (
	"log"

	"backend/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	db, err := DBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	firestoreClient, err := InitFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Api is running!"})
	})

	controller.AuthenticateController(router, db, firestoreClient)

	router.Run()
}
