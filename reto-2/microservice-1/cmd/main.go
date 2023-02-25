package main

import (
	files "asalaza5-st0263/reto-2/microservice-1/internal/proto/files"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type FileServiceServer struct {
	files.UnimplementedFileServiceServer
	directory string
}

func (f *FileServiceServer) ListFiles(ctx context.Context, _ *files.FileListRequest) (*files.FileListResponse, error) {
	dirFiles := listFiles(f.directory)
	r := files.FileListResponse{Files: dirFiles}
	return &r, nil
}

func (f *FileServiceServer) SearchFiles(ctx context.Context, req *files.FileSearchRequest) (*files.FileSearchResponse, error) {
	filteredFiles := searchFiles(f.directory, req.Query)
	r := files.FileSearchResponse{Files: filteredFiles}
	return &r, nil
}

func main() {
	port := os.Getenv("MICROSERVICE_ONE_PORT")
	ip := os.Getenv("MICROSERVICE_ONE_IP")
	directory := os.Getenv("MICROSERVICE_ONE_DIRECTORY")

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	filesServiceServer := &FileServiceServer{directory: directory}
	files.RegisterFileServiceServer(server, filesServiceServer)
	reflection.Register(server)
	log.Printf("Listening on port: %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func listFiles(directory string) []string {
	var dirFiles []string
	dirEntries, _ := os.ReadDir(directory)
	for _, file := range dirEntries {
		dirFiles = append(dirFiles, file.Name())
	}
	return dirFiles
}

func searchFiles(directory, name string) []string {
	dirFiles := listFiles(directory)
	var ret []string
	if name == "*" {
		return dirFiles
	}
	for _, file := range dirFiles {
		if strings.Contains(file, name) {
			ret = append(ret, file)
		}
	}
	return ret
}
