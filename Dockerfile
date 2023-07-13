FROM docker.io/golang:latest

WORKDIR /go/build
COPY . .
RUN go build -o /usr/local/bin/app

EXPOSE 8123
ENV ASSET_PATH=/srv/app
ENTRYPOINT ["/usr/local/bin/app"]