FROM golang:1.19-alpine

COPY . /code

WORKDIR /code
RUN go install

ENTRYPOINT [ "topo" ]