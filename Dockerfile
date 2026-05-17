FROM golang AS builder
WORKDIR /app
COPY . .                  
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .     






    