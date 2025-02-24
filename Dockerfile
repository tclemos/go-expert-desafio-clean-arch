FROM golang:1.23.5 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o grpc ./cmd/grpc/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rest ./cmd/rest/

FROM golang:1.23.5 as grpc
WORKDIR /app
COPY --from=build /app/grpc .
COPY --from=build /app/cmd/grpc/.env .
ENTRYPOINT ["./grpc"]

FROM scratch as rest
WORKDIR /app
COPY --from=build /app/rest .
COPY --from=build /app/cmd/rest/.env .
ENTRYPOINT ["./rest"]
