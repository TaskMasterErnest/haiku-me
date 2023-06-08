# # get the base image, no vulnerabilities here
# FROM golang:1.20.4-alpine3.18 AS builder

# # create a working directory
# WORKDIR /app

# # # copy over the mod and sum files to the workdir
# # COPY go.mod go.sum ./
# # # install the go.mod 
# # RUN go mod download

# # copy all except the database dir
# COPY . .

# # build the go application
# # RUN go build -o /run/web ./cmd/web/ && cp -r ./tls /run/

# EXPOSE 4000

# # run the binary
# CMD ["go","run","./cmd/web"]

FROM golang:1.20.4-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/web/
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/tls ./tls
EXPOSE 4000
CMD [ "./main" ]