# Sử dụng Go image chính thức (alpine nhẹ hơn)
FROM golang:1.23-alpine

# Tạo thư mục làm việc
WORKDIR /app

# Copy file go vào container
COPY . .

# Tải dependency và build ứng dụng
RUN go mod tidy && go build -o main .

# Mở cổng 8001
EXPOSE 8001

# Lệnh chạy khi container khởi động
CMD ["./main"]
