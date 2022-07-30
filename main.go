package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edubarbieri/julius/auth"
	"github.com/edubarbieri/julius/config"
	"github.com/edubarbieri/julius/repository"
	"github.com/edubarbieri/julius/service"
	"github.com/edubarbieri/julius/web"
	"github.com/gin-gonic/gin"
)

func main() {
	nfeRepository, err := repository.NewPgNfeRepository(config.PostgresURL)
	if err != nil {
		log.Fatalf("error creating PgNfeRepository %v", err)
	}
	nfeService := service.NewNfeService(nfeRepository)
	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.Use(auth.JwtTokenCheck)
	{
		v1Api := web.NewV1Api(nfeService)
		v1Api.SetupRouters(v1)
	}

	router.StaticFS("/", http.Dir("./frontend"))

	log.Println("Starting server in port", config.HttpPort)
	router.Run(fmt.Sprintf(":%d", config.HttpPort))
}
