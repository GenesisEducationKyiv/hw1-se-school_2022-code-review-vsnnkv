FROM golang:1.18-alpine
EXPOSE 8181

COPY go.mod ./
COPY go.sum ./
ENV GOPATH=/

RUN go mod download

COPY .. ./

RUN go build -o customers ./main.go


CMD [ "./customers"]