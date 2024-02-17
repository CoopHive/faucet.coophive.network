FROM node:20-alpine as frontend

WORKDIR /app

COPY web/package.json web/pnpm-lock.yaml ./
COPY web ./

RUN npm install -g pnpm
RUN pnpm install

RUN npm run build

FROM golang:1.21 as backend

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
RUN mkdir -p ./bin

#RUN apk add --no-cache gcc musl-dev linux-headers

RUN go install github.com/goreleaser/goreleaser@latest

COPY go.* ./
RUN go mod download

COPY . .

COPY --from=frontend /app/dist web/dist


RUN goreleaser build --single-target --clean -o ./bin/faucet --snapshot

#RUN ./bin/faucet version

FROM alpine:latest

ENV PORT=8080

WORKDIR /bin

RUN apk add --no-cache ca-certificates

COPY --from=backend /app/bin/faucet /bin/faucet

EXPOSE 8080

ENTRYPOINT ["/bin/faucet"]

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"