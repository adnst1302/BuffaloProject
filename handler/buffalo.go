package handler

import (
	"log"
	"net/http"

	"BuffaloProject/buffalo"

	"github.com/gin-gonic/gin"
)

func BuffaloDetail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"Page": "Detail Buffalo",
		"Data": id,
	})
}

func BuffaloAdd(c *gin.Context) {
	var DataInput buffalo.Input

	err := c.ShouldBindJSON(&DataInput)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"name":       DataInput.Name,
		"family":     DataInput.Family,
		"population": DataInput.Population,
	})

}
