FROM golang:1.13.4-alpine
COPY . /go/src/wru-backend-new-merchant-on-boarding/
WORKDIR "/go/src/wru-backend-new-merchant-on-boarding/app"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM alpine:latest
WORKDIR /root
COPY --from=0 /go/src/wru-backend-new-merchant-on-boarding/swaggerui ./swaggerui
COPY --from=0 "/go/src/wru-backend-new-merchant-on-boarding/app/app" .
CMD ["./app","-file=dev.json"]
#testing

