package main

import (
	"crypto/tls"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	port := "50051"

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	creds, err := credentials.NewServerTLSFromFile("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(creds))
	// ... register gRPC services ...

	// imageServiceServer := &imageService.ImageServiceServer{}
	// messages.RegisterImageServiceServer(server, imageServiceServer)
	// if err := server.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }

	reflection.Register(server)

	log.Printf("Listening on port: %v\n", port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
