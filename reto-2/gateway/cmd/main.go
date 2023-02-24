package main

import (
	"os"
	"reto-2/gateway/cmd/routes"
	"reto-2/gateway/internal/client"
	"reto-2/gateway/internal/rabbitmq"

	"github.com/gin-gonic/gin"
)

func main() {

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

	// amqpConn, err := amqp.Dial(rabbitmqURL)
	// if err != nil {
	// 	panic(err)
	// }
	// defer amqpConn.Close()
	rabbitmqService, err := rabbitmq.Setup(rabbitmqURL)
	if err != nil {
		panic(err)
	}
	defer rabbitmqService.Conn.Close()

	router := routes.NewRouter(eng, grpcConn, rabbitmqService)
	router.MapRoutes()

	if err := eng.Run(":" + port); err != nil {
		panic(err)
	}

}
