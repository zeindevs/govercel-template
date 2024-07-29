package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func Handler(w http.ResponseWriter, r *http.Request) {
  app.ServeHTTP(w, r)
}

func init() {
  app = gin.New()

  app.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, map[string]any{
      "msg": "Hello Bro! from Gin",
    })
  })
}


