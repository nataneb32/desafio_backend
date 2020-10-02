FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

#Command to run the executable
ENTRYPOINT CompileDaemon --build="go build cmd/web/main.go" --command=./main


