package routes

import (
	"asalaza5-st0263/reto-2/gateway/cmd/handler"
	"asalaza5-st0263/reto-2/gateway/internal/rabbitmq"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng             *gin.Engine
	grpcConn        *grpc.ClientConn
	rabbitmqService *rabbitmq.RabbitMQService
}

func NewRouter(eng *gin.Engine, grpcConn *grpc.ClientConn, rabbitmqService *rabbitmq.RabbitMQService) Router {

	return &router{eng: eng, grpcConn: grpcConn, rabbitmqService: rabbitmqService}
}

func (r *router) MapRoutes() {
	r.buildHandlerRoutes()
}

func (r *router) buildHandlerRoutes() {
	handler := handler.NewHanlder(r.grpcConn, r.rabbitmqService)
	r.eng.GET("/list", handler.ListFiles())
	r.eng.GET("/search", handler.SearchFiles())
}
