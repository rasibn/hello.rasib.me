FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build  -ldflags="-s -w" -o ./build/release ./server
CMD ["./build/release"]
