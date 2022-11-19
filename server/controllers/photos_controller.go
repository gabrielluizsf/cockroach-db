package controllers

import (
	"strconv"

	"github.com/GabrielLuizSF/cockroach-db/database"
	"github.com/GabrielLuizSF/cockroach-db/database/models"
	"github.com/GabrielLuizSF/cockroach-db/server/utils/errors/client"
	"github.com/gin-gonic/gin"
)

func ShowAllPhotos(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Photos
	err := db.Find(&p).Error
	client.PrintClientError(err, "cannot find photos by id: ", c)
	c.JSON(200, p)
}

func ShowPhoto(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	client.PrintClientError(err, "ID has to be integer", c)
	db := database.GetDatabase()
	var p models.Photos
	err = db.First(&p, newid).Error
	client.PrintClientError(err, "cannot find photos by id: ", c)

	c.JSON(200, p)
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Photos

	err := c.ShouldBindJSON(&p)
	client.PrintClientError(err, "cannot bind JSON: ", c)
	err = db.Create(&p).Error
	client.PrintClientError(err, "cannot create photos: ", c)
	c.JSON(200, p)
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	client.PrintClientError(err, "ID has to be integer", c)

	db := database.GetDatabase()

	err = db.Delete(&models.Photos{}, newid).Error
	client.PrintClientError(err, "cannot delete photos: ", c)

	c.Status(204)
}

func EditPhoto(c *gin.Context) {
	db := database.GetDatabase()
	var p models.Photos

	err := c.ShouldBindJSON(&p)
	client.PrintClientError(err, "cannot bind JSON: ", c)
	err = db.Save(&p).Error
	client.PrintClientError(err, "cannot create photos: ", c)

	c.JSON(200, p)
}
