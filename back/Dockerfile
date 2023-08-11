FROM golang:1.21.0

WORKDIR /back
COPY ./ /back

RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080

CMD ["go", "run", "main.go"]