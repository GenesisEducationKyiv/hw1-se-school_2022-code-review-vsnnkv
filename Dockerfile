FROM golang:1.18-alpine
EXPOSE 8080

COPY go.mod ./
COPY go.sum ./
ENV GOPATH=/

RUN go mod download

COPY ./ ./

RUN go build -o btc-app ./cmd/main.go


CMD [ "./btc-app" ]