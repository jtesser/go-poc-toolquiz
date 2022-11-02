####### Start from the latest golang base image ###############
FROM golang:1.19.3-buster as builder
LABEL maintainer="Jason Tesser <jasontesser@gmail.com>"
WORKDIR /app
COPY ./ ./
RUN go mod download
# Build the Go app.  This below command builds Go statically linked so that all C libs etc... are included
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/main .

######## Image to run Unit Tests #######
#### can use later to add Unit tests ########
#FROM builder as unit-test
#COPY docker/entry-utest.sh /usr/local/bin
#RUN chmod 755 /usr/local/bin/entry-utest.sh
#ENTRYPOINT ["/usr/local/bin/entry-utest.sh"]

######## Start a new stage from a small image #######
FROM alpine:latest as prod
# Add root certs
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the builder stage
COPY --from=builder /usr/local/bin/main .
EXPOSE 8080
COPY docker/entrypoint.sh /usr/local/bin
RUN chmod 755 /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
