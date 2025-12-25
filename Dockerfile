FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service-order ./cmd

FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=build /app/service-order .

EXPOSE 5001
CMD ["/app/service-b"]



