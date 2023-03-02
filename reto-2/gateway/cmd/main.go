package main

import (
	"asalaza5-st0263/reto-2/gateway/cmd/routes"
	"asalaza5-st0263/reto-2/gateway/internal/client"
	"asalaza5-st0263/reto-2/gateway/internal/rabbitmq"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	time.Sleep(60 * time.Second)
	port := os.Getenv("GATEWAY_PORT")
	// ip := os.Getenv("GATEWAY_IP")
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	// address := fmt.Sprintf("%v:%v", ip, port)

	eng := gin.Default()

	grpcConn, err := client.NewGrpcClient()

	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	rabbitmqService, err := rabbitmq.Setup(rabbitmqURL)
	if err != nil {
		panic(err)
	}
	defer rabbitmqService.Connection.Close()

	router := routes.NewRouter(eng, grpcConn, rabbitmqService)
	router.MapRoutes()

	if err := eng.Run(":" + port); err != nil {
		panic(err)
	}

}
