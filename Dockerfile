# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:latest AS builder
RUN mkdir /go/src/mtv
WORKDIR /go/src/mtv

COPY movgolib.go .
COPY mtvlib.go .
COPY mtvserver.go .
COPY thumbgolib.go .
COPY tvgolib.go .
COPY typeslib.go .
COPY go.mod .
COPY go.sum .

RUN export GOPATH=/go/src/mtv
RUN go get -v /go/src/mtv
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/mtv

# FROM arm32v6/alpine:latest
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/mtv/main .
RUN \
  mkdir ./static && \
  chmod -R +rwx ./static && \
  mkdir ./fsData && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/mtvServer.log && \
  echo "Creating log file" > ./logs/mtvTV.log && \
  echo "Creating log file" > ./logs/mtvMOV.log && \
  chmod -R +rwx ./logs/mtvServer.log && \
  chmod -R +rwx ./logs/mtvTV.log && \
  chmod -R +rwx ./logs/mtvMOV.log
  
CMD ["./main"]
STOPSIGNAL SIGInt
