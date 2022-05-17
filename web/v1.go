package web

import (
	"net/http"

	"github.com/edubarbieri/julius/service"
	"github.com/gin-gonic/gin"
)

type V1Api struct {
	nfeService service.NfeService
}

func NewV1Api(nfeService service.NfeService) V1Api {
	return V1Api{
		nfeService: nfeService,
	}
}

func (v *V1Api) SetupRouters(route gin.IRoutes) {
	route.POST("/nfe", v.postNfe)
}

func (v *V1Api) postNfe(ctx *gin.Context) {
	request := postNfeRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := v.nfeService.SaveNfe(ctx, request.Url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "nfe importada com sucesso"})
}
