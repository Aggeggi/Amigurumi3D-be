package main

import (
	"os"
	"time"

	"github.com/Aggeggi/Amigurumi3D-be/database"
	_ "github.com/Aggeggi/Amigurumi3D-be/docs"
	"github.com/Aggeggi/Amigurumi3D-be/middlewares"
	"github.com/Aggeggi/Amigurumi3D-be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowedOrigins
		},
		MaxAge: 12 * time.Hour,
	}))
	cookieSecret := os.Getenv("COOKIE_SECRET")
	cookieName := os.Getenv("COOKIE_NAME")
	store := cookie.NewStore([]byte(cookieSecret))
	r.Use(sessions.Sessions(cookieName, store))

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.AuthRequired())
	{

		v1.GET("/ping", routes.Ping)
		v1.POST("/amigurumi", routes.PostAmigurumi)
		v1.GET("/amigurumi", routes.GetListAmigurumi)
		v1.GET("/amigurumi/:id", routes.GetAmigurumi)
		v1.POST("/logout", routes.Logout)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/v1/login", routes.Login)
	r.POST("/api/v1/signin", routes.SignIn)
	port := os.Getenv("PORT")
	r.Run(port)
}
