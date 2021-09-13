FROM golang:1.14

EXPOSE 1331

COPY . ${GOPATH}/src/db
WORKDIR ${GOPATH}/src/db

RUN go build
CMD ["./db"]