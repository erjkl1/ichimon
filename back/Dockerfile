FROM golang:1.21.0

WORKDIR /back
COPY ./ /back

RUN go install github.com/cosmtrek/air@v1.44.0

RUN go mod tidy \
  && go build

# DEVはあとで消す
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64\
  GO_ENV=dev 
EXPOSE 8080

CMD ["go", "run", "main.go"]