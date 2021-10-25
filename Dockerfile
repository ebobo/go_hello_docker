FROM golang:1.17.2-alpine3.14

WORKDIR /app
COPY . ./

RUN go mod download
RUN go build -o /hello

# Start app
CMD ["/hello","greeting","-n qi" ]