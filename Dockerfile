# Use an official Go runtime as the parent image
FROM golang:1.19-alpine as builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o acmverifier .

# Use a smaller image to run our application
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /go/src/app/acmverifier .

# Command to run the application
ENTRYPOINT ["./acmverifier"]

