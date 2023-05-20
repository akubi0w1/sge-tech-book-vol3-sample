FROM --platform=linux/amd64 golang:1.19-alpine as build

WORKDIR /api
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal

RUN GOOS=linux GOARCH=amd64 go build -o /api/bin/api ./cmd/api/main.go

FROM --platform=linux/amd64 alpine
COPY --from=build /api/bin/api /api

CMD ["/api"]