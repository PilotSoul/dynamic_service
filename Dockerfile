FROM debian:buster-slim

RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential

ENV GOLANG_VERSION 1.18.3
RUN curl -fsSL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# RUN export PATH=$PATH:$(go env GOPATH)/bin
RUN export GOPATH="$HOME/go"
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker

CMD ["/docker"]