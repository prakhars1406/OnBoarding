FROM golang:1.13.4-alpine
COPY . /go/src/recommendation-rendering-service/
WORKDIR "/go/src/wru-backend-new-merchant-on-boarding/app"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM alpine:latest
WORKDIR /root
COPY ./devops/aws_prod/prod.json .
COPY --from=0 /go/src/recommendation-rendering-service/swaggerui ./swaggerui
COPY --from=0 "/go/src/wru-backend-new-merchant-on-boarding/app/app" .
CMD ["./app","-file=prod.json"]
#testing
