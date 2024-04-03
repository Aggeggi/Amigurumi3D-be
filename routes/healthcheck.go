package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

// PingExample godoc
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/ping [get]
func Ping(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
}