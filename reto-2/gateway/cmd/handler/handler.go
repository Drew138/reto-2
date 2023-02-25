package handler

import (
	files "asalaza5-st0263/reto-2/gateway/internal/proto/files"
	"asalaza5-st0263/reto-2/gateway/internal/rabbitmq"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Handler struct {
	client          files.FileServiceClient
	rabbitmqService *rabbitmq.RabbitMQService
}

func NewHanlder(conn *grpc.ClientConn, rabbitmqService *rabbitmq.RabbitMQService) *Handler {
	fileServiceClient := files.NewFileServiceClient(conn)
	return &Handler{client: fileServiceClient, rabbitmqService: rabbitmqService}
}

var count = 0

func (h *Handler) ListFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		count++
		count %= 2
		if count == 0 {
			response, err := h.client.ListFiles(context.Background(), &files.FileListRequest{})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"response": response.Files, "type": "rpc"})
			}
		} else {
			queueName := uuid.New().String()
			request := struct {
				Name      string `json:"name"`
				Query     string `json:"query"`
				QueueName string `json:"queue_name"`
			}{Name: "list", QueueName: queueName}

			body := &bytes.Buffer{}
			enc := json.NewEncoder(body)
			enc.Encode(request)
			h.rabbitmqService.Publish(body.Bytes(), queueName)
			channel, err := h.rabbitmqService.Receive(queueName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			msg := <-channel
			_, err = h.rabbitmqService.Channel.QueueDelete(queueName, false, false, false)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			var response []string

			err = json.Unmarshal(msg.Body, &response)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"response": response, "type": "mom"})
		}
	}
}

func (h *Handler) SearchFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		count++
		count %= 2
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
			queueName := uuid.New().String()
			request := struct {
				Name      string `json:"name"`
				Query     string `json:"query"`
				QueueName string `json:"queue_name"`
			}{Name: "search", Query: query, QueueName: queueName}

			body := &bytes.Buffer{}
			enc := json.NewEncoder(body)
			enc.Encode(request)
			h.rabbitmqService.Publish(body.Bytes(), queueName)
			channel, err := h.rabbitmqService.Receive(queueName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			msg := <-channel
			_, err = h.rabbitmqService.Channel.QueueDelete(queueName, false, false, false)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			var response []string

			err = json.Unmarshal(msg.Body, &response)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"response": response, "type": "mom"})
		}
	}
}
