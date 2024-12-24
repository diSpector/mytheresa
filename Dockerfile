FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum main.go ./
RUN go mod download

COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /mytheresa

FROM build-stage AS test-stage
RUN go test -v ./...

FROM alpine:latest

WORKDIR /
COPY --from=build-stage /mytheresa /mytheresa

EXPOSE 8082

CMD ["/mytheresa"]