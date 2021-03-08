FROM golang:1.16

ARG PORT

WORKDIR /go/src/github.com/karamaru-alpha/grpc-sample

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE ${PORT}

CMD ["fresh"]
