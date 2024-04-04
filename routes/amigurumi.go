package routes

import (
	"github.com/Aggeggi/Amigurumi3D-be/requests"
	"github.com/Aggeggi/Amigurumi3D-be/entities"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

//	@BasePath	/api/v1

// PingExample godoc
//	@Summary	Insert an amigurumi in db
//	@Schemes
//	@Description	Add a new amigurumi
//	@Tags			amigurumi
//	@Accept			json
//	@Produce		json
// @Param amigurumi body requests.PostAmigurumiRequest true "amigurumi info"
//	@Success		200	{string}	Amigurumi name
//	@Router			/amigurumi [post]
func PostAmigurumi(context *gin.Context)  {
	body:=requests.PostAmigurumiRequest{}
      if err:=context.BindJSON(&body);err!=nil{
         context.AbortWithError(http.StatusBadRequest,err)
         return
      }
	amigurumi := entities.NewAmigurumiPattern(body.Name)
    if err := mgm.Coll(amigurumi).Create(amigurumi);err!=nil{
         context.AbortWithError(http.StatusBadRequest,err)
         return
      }
}