FROM golang:1.12 AS build

WORKDIR /src/
COPY main.go /src/
COPY go.mod go.sum /src/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/webserver

FROM alpine:latest
COPY --from=build /bin/webserver /bin/webserver
COPY health /bin/
WORKDIR /bin/
ENTRYPOINT ["/bin/webserver"]