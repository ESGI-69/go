FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY . .

WORKDIR /usr/src/app/src
CMD swag init && cd .. && go run src/main.go
