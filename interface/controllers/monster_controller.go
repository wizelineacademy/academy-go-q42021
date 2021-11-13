package controllers

import (
	"net/http"

	"github.com/Le-MaliX/ACADEMY-GO-Q42021/services"
	"github.com/gin-gonic/gin"
)

func GetAllMonsters(c *gin.Context) {
	monsters, err := services.GetAllMonsters()
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	c.JSON(http.StatusOK, monsters)
}

func GetMonsterById(c *gin.Context) {
	id := c.Param("id")
	monster, err := services.GetMonsterById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	c.JSON(http.StatusOK, monster)
}
