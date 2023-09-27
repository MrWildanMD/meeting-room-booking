# Use an official Golang runtime as a parent image
FROM golang:alpine

# Add maintaner info
LABEL maintaner="Wildan Mauluddana"

# Update the container dependency and add git
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

# Copy the local package files to the container's workspace
COPY . .

# Run go mod tidy for installing/removing unused package
RUN go mod tidy