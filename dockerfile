# Use an official Golang runtime as the base image
FROM golang:alpine

RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module dependency files
COPY go.mod go.sum ./

# Copy the source code into the container
COPY . .

# Download the Go dependencies
RUN go mod download


# Build the Go application
# ubuntu
RUN go build -o myapp .

# windows
# RUN go build -o myapp.exe .
# RUN go build .

RUN ls

# Set the entry point command for the container
# CMD ["ls"]
CMD ["./myapp"]
