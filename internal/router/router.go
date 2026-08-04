package router

import (
	"net/http"

	"github.com/Golukpal/url-shortener/internal/app"

	"github.com/gin-gonic/gin"
)

func Setup(a *app.App) *gin.Engine {

	r := gin.New()

	r.GET("/health", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})

	return r
}
