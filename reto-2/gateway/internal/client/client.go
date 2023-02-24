package client

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
)

func NewGrpcClient() (*grpc.ClientConn, error) {
	ip, port := os.Getenv("MICROSERVICE_ONE_IP"), os.Getenv("MICROSERVICE_ONE_PORT")
	address := fmt.Sprintf("%v:%v", ip, port)
	return grpc.Dial(address, grpc.WithInsecure())
}
