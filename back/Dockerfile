FROM golang:1.19-alpine

COPY . /code

WORKDIR /code
RUN go install
RUN mkdir /tmp/local_modules

ENTRYPOINT [ "topo" ]