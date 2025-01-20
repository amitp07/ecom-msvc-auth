FROM golang:1.23.4-alpine3.21
WORKDIR /app  
COPY go.mod ./ 
RUN go mod download
COPY . . 
RUN go build -o main ./cmd/api/*
EXPOSE 5000
CMD [ "./main" ]


