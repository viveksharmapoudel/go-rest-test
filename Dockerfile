FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /rest-api-gin-gorm-sql
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
ENV PORT 8000
RUN go build
CMD ["./rest-api-test"]