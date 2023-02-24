package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	files "reto-2/gateway/internal/proto/files"
	"reto-2/gateway/internal/rabbitmq"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Handler struct {
	client          files.FileServiceClient
	rabbitmqService *rabbitmq.ServiceAMQP
}

func NewHanlder(conn *grpc.ClientConn, rabbitmqService *rabbitmq.ServiceAMQP) *Handler {
	fileServiceClient := files.NewFileServiceClient(conn)
	return &Handler{client: fileServiceClient, rabbitmqService: rabbitmqService}
}

var count = 0

func (h *Handler) ListFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		if count == 0 {
			response, err := h.client.ListFiles(context.Background(), &files.FileListRequest{})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"response": response.Files, "type": "rpc"})
			}
		} else {

			request := struct {
				Name  string `json:"name"`
				Query string `json:"query"`
			}{Name: "list"}

			body := &bytes.Buffer{}
			enc := json.NewEncoder(body)
			enc.Encode(request)
			h.rabbitmqService.Publish(body.Bytes())
			response := <-h.rabbitmqService.C

			c.JSON(http.StatusOK, gin.H{"response": response, "type": "mom"})
		}
		count++
		count %= 2
	}
}

func (h *Handler) SearchFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		if count == 0 {
			query := c.Query("query")
			response, err := h.client.SearchFiles(context.Background(), &files.FileSearchRequest{Query: query})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"response": response.Files, "type": "rpc"})
			}
		} else {
			query := c.Query("query")
			request := struct {
				Name  string `json:"name"`
				Query string `json:"query"`
			}{Name: "search", Query: query}

			body := &bytes.Buffer{}
			enc := json.NewEncoder(body)
			enc.Encode(request)
			h.rabbitmqService.Publish(body.Bytes())
			response := <-h.rabbitmqService.C

			c.JSON(http.StatusOK, gin.H{"response": response, "type": "mom"})
		}
		count++
		count %= 2
	}
}

func listFiles(directory string) []string {
	var files []string

	dirEntries, _ := os.ReadDir(directory)
	for _, file := range dirEntries {
		files = append(files, file.Name())
	}
	return files
}

func searchFiles(directory, name string) []string {
	files := listFiles(directory)
	var ret []string
	if name == "*" {
		return files
	}
	for _, file := range files {
		if strings.Contains(file, name) {
			ret = append(ret, file)
		}
	}
	return ret
}
