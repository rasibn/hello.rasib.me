FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build  -ldflags="-s -w" -o ./build/release .
CMD ["./build/release"]
