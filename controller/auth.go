package controller

import (
	"backend/model"
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthenticateController(router *gin.Engine, db *gorm.DB, firestoreClient *firestore.Client) {
	routes := router.Group("/auth")
	{
		routes.GET("/", func(c *gin.Context) {
			GetAllProduct(c, db, firestoreClient)
		})
		routes.GET("/firebase", func(c *gin.Context) {
			GetAllProductFirebase(c, db, firestoreClient)
		})
	}
}

func GetAllProduct(c *gin.Context, db *gorm.DB, firestoreClient *firestore.Client) {
	var user []model.User
	result := db.Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllProductFirebase(c *gin.Context, db *gorm.DB, firestoreClient *firestore.Client) {
	// ดึงข้อมูลทั้งหมดจาก Collection "users" ใน Firestore
	ctx := context.Background()
	iter := firestoreClient.Collection("usersLogin").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get documents: %v", err)})
		return
	}

	// สร้าง slice เพื่อเก็บข้อมูลทั้งหมด
	var users []map[string]interface{}
	for _, doc := range docs {
		data := doc.Data()
		users = append(users, data)
	}

	// ส่งข้อมูลทั้งหมดกลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{"data": users})
}
