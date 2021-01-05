FROM golang:1.15 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download

RUN ls /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o abap-mock-in-go /app

RUN ls /app/

FROM scratch
WORKDIR /app
COPY --from=builder /app/abap-mock-in-go /app/

EXPOSE 8000
ENTRYPOINT ["/app/abap-mock-in-go"]