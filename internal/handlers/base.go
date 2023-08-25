package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func LoadHomePage(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	}

