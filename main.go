package main

import (
	"fmt"
	"log"
	"net/http"

	"BuffaloProject/buffalo"
	"BuffaloProject/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "root:3ncoding4sc11A@@tcp(127.0.0.1:3306)/dbs_buffalo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(buffalo.Buffalo{})
	fmt.Println("db berhasil connect")


	// buffalo := buffalo.Buffalo{}
	// buffalo.Name = "Alfi Ananda"
	// buffalo.Family = "Aogra"
	// buffalo.Population = 140

	// db.Create(&buffalo)

	router.GET("/", HomePage)

	v1 := router.Group("/v1")
	v1.GET("/buffalo/:id", handler.BuffaloDetail)
	v1.POST("/buffalo/add", handler.BuffaloAdd)

	router.Run(":7001")
}

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Project": "Fresh Buffalo",
	})
}
