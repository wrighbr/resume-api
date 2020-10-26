FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/resume-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/http-swagger
RUN go get -u github.com/alecthomas/template

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Building swagger
RUN swag init

# Build the Go app
RUN go build -o ./out/resume-api .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/resume-api/out/resume-api /app/resume-api

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/resume-api"]
