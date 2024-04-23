# Gunakan image Go sebagai base image
FROM golang:latest

# Set working directory di dalam container
WORKDIR /app

# Copy seluruh file backend ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Command default untuk menjalankan aplikasi backend
CMD ["./main"]
