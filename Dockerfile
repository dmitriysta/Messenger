
FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /messenger

EXPOSE 8080

CMD [ "/messenger" ]