FROM golang:1.24-alpine AS build

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app src/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /

CMD ["/app"]