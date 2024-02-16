FROM node:20-alpine as frontend

WORKDIR /app

COPY web/package.json web/pnpm-lock.yaml ./
COPY web ./

RUN npm install -g pnpm
RUN pnpm install

RUN npm run build

FROM golang:1.21-alpine as backend

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
RUN mkdir -p ./bin

RUN apk add --no-cache gcc musl-dev linux-headers

RUN go install github.com/goreleaser/goreleaser@latest

COPY go.* ./
RUN go get

COPY . .

COPY --from=frontend /app/dist web/dist


RUN goreleaser build --single-target --clean -o ./bin/coophive-faucet --snapshot

#RUN ./bin/coophive-faucet version

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=backend /app/coophive-faucet /app/coophive-faucet

EXPOSE 8080

ENTRYPOINT ["/app/coophive-faucet"]

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"