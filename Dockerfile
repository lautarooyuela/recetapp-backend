FROM golang:latest
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main .
EXPOSE 4000
CMD ["./main"]