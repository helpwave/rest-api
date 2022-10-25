FROM golang:1.19 AS builder

WORKDIR /app
COPY . /app

ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.Version=${VERSION}" -a -o helpwave .

FROM alpine AS production

LABEL maintainer="development@helpwave.de"

WORKDIR /app

COPY --from=builder /app/helpwave .
RUN chmod +x helpwave

ENV PORT 80
ENV GIN_MODE release
ENV LOG_LEVEL info
ENV POSTGRES_PORT 5432

EXPOSE 80

CMD ["./helpwave"]
