# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o threatintelapi main.go

# Run stage
FROM alpine:latest
WORKDIR /app
ENV PORT=8000
ENV TZ=Europe/Brussels
COPY --from=builder /app/threatintelapi ./threatintelapi
COPY ./resources ./resources
COPY ./openapi.yaml ./openapi.yaml
COPY ./docs ./docs
copy ./resources/openphish ./resources/openphish
RUN apk add --no-cache git busybox-openrc openrc
RUN mkdir -p /var/log && printf '0 2 * * * git -C /app/resources/openphish pull --ff-only --quiet' > /etc/crontabs/root
EXPOSE 8000
CMD ["/bin/sh","-c","crond -b -l 8; ./threatintelapi"]