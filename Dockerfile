# Start with the Go base image
FROM golang:1.23.1

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project (scraper) into the container
COPY . .

# Ensure the binary is built for the Linux/AMD64 architecture
ENV GOOS=linux
ENV GOARCH=amd64

# Download necessary Go modules
RUN go mod download

# Build the Go application for Linux/AMD64
RUN go build -o /app/main ./cmd/main.go || { echo 'Build failed!'; exit 1; }

# Debugging: List the contents of the /app directory
RUN echo "Listing /app directory after build:" && ls -l /app

# Ensure the binary is executable
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ["/app/main"]
