FROM golang:1.16
ARG DOMAIN
ENV DOMAIN=${DOMAIN:-localhost}

WORKDIR /go/src/backend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD backend $DOMAIN