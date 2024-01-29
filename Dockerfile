FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o app

# Expose port 8080 for the application
EXPOSE 9090

# Set the default command to run when the container starts
CMD ["./app"]