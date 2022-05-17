package main

import (
	"log"
	"os"

	"github.com/edubarbieri/julius/repository"
	"github.com/edubarbieri/julius/service"
	"github.com/edubarbieri/julius/web"
	"github.com/gin-gonic/gin"
)

func main() {
	nfeRepository, err := repository.NewPgNfeRepository(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("error creating PgNfeRepository %v", err)
	}
	nfeService := service.NewNfeService(nfeRepository)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1Api := web.NewV1Api(nfeService)
		v1Api.SetupRouters(v1)
	}
	log.Println("Starting server in port 8081")
	router.Run(":8081")
}
