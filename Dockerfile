
FROM golang:latest


WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o main .


EXPOSE 3777

# Command to run the executable
CMD ["./main"]
