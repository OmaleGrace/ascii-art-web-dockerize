# Stage 1: builder
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

# Stage 2: final image
FROM alpine:

# Metadata
LABEL maintainer="oomale"
LABEL project="ascii-art-web-dockerize"
LABEL description="ASCII Art Web application built with Go"

WORKDIR /app
COPY --from=builder /app/myapp .

# Copy any required assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banner ./banner

EXPOSE 8080

CMD ["./myapp"]
