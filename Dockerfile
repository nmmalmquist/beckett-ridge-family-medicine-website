FROM  --platform=linux/arm64 golang:1.22.3 AS build
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
RUN go build -v -o app .

FROM scratch
COPY --from=build /go/src/app/app /go/bin/app
EXPOSE 8000
ENTRYPOINT ["/go/bin/app"]