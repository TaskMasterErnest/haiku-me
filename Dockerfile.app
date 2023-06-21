FROM golang:1.20.4-alpine3.18 AS builder
# create a working directory
WORKDIR /app
# copy over the mod and sum files to the workdir
COPY go.mod go.sum ./
# install the go.mod 
RUN go mod download
# copy all except the database/container
COPY . .
# build the go application
RUN CGO_ENABLED=0 go build -o main ./cmd/web/

# start another container
FROM alpine:3.18
# set the working directory
WORKDIR /app
# copy the relevant files from previous container
COPY --from=builder /app/main .
COPY --from=builder /app/tls ./tls
# expose a particular port
EXPOSE 4000
# run the binary
CMD [ "./main" ]