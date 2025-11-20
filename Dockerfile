# Multi-stage build for optimized image size
FROM golang:1.23-alpine AS backend-builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory for backend
WORKDIR /app/backend

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy backend source code
COPY backend/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage - minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Set timezone (adjust as needed)
ENV TZ=Europe/Paris

WORKDIR /app

# Copy the built binary from builder
COPY --from=backend-builder /app/backend/main .

# Copy frontend files
COPY frontend/ ./frontend/

# Copy backend data file
COPY backend/date_link.txt ./

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
