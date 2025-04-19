FROM golang:1.24.2-bookworm AS base

FROM base AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest

COPY . .
CMD [ "air", "-c", ".air.toml" ]