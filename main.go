package main

import (
	_ "github.com/Aggeggi/Amigurumi3D-be/docs"
	"github.com/Aggeggi/Amigurumi3D-be/routes"
    "github.com/Aggeggi/Amigurumi3D-be/database"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
//	@title			Amigurumi 3D
//	@version		1.0
//	@description	This is a server for the amigurumi 3D app

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	database.InitDb()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{

		v1.GET("/ping", routes.Ping)
		v1.POST("/amigurumi", routes.PostAmigurumi)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
