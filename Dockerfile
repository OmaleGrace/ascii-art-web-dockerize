# Stage 1: builder
FROM golang:1.22.2 AS builder
WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp . 

# Stage 2: final image
FROM alpine:latest

# Metadata
LABEL maintainer="oomale"
LABEL project="ascii-art-web-dockerize"
LABEL description="ASCII Art Web application built with Go"

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/ascii ./ascii
COPY --from=builder /app/banner ./banner
COPY --from=builder /app/handlers ./handlers

EXPOSE 8080

CMD ["./myapp"]
