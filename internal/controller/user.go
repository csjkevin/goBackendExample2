package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"goBackendExample/internal/model"
	"goBackendExample/internal/service"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success"})
}

func GetAllUser(c *gin.Context) {
	result, err := service.GetAllUser()
	if err != nil {
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success", "result": result})
}

func GetUserById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Panic(err)
		return
	}

	result, err := service.GetUserById(id)
	if err != nil {
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success", "result": result})
}

func UpdateUser(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Panic(err)
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateUser(id, user)
	if err != nil {
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success"})
}