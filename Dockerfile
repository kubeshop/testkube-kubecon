# syntax=docker/dockerfile:1
FROM golang:1.21
WORKDIR /svc/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN mkdir /tmp/data
WORKDIR /root/
COPY --from=0 /svc/app /bin/runner
CMD ["/bin/runner"]  