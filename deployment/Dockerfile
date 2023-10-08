FROM golang:1.21.0-alpine as base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOD=linux go build -o /exe

FROM gcr.io/distroless/static-debian11

COPY --from=base /exe .

CMD ["/exe"]
