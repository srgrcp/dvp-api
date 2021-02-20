FROM golang:1.15-alpine
WORKDIR /go/src/dvp-api
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o dvp-api main.go
EXPOSE 3700
CMD [ "./dvp-api" ]
