 #Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:latest
# Add Maintainer Info
LABEL maintainer="Tamil Vanan<tamilhce@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

ADD ../server /app/
# Copy go mod and sum files
#COPY ../server/go.mod ../server/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
#COPY ../server .

# Build the Go app
# RUN go build -o pkg/server .

# Expose port 5000 to the outside world
# EXPOSE 5000

# Command to run the executable
# CMD ["./server"]