.PHONY: install run proto

run:
	@go docker-compose up

proto:
	@protoc --go_out=./gateway/internal --go-grpc_out=./gateway/internal ./files.proto
	@protoc --go_out=./microservice-1/internal --go-grpc_out=./microservice-1/internal ./files.proto

install:
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo Installation complete.
	@echo Remember to add \`\export \GOPATH\=$\HOME/go\` 
	@echo and \`\export \PATH\=$\PATH:$GOPATH/bin\`
	@echo to your .bashrc file.

rabbitmq:
	@docker run rabbitmq:3.8-management-alpine -d -p 15672:15672 -p 5672:5672 \
		-v ${PWD}/rabbitmq/rabbitmq-definitions.json:/etc/rabbitmq/definitions.json

