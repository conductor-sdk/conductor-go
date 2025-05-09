FROM golang:1.19 as build
RUN mkdir /package
COPY /sdk /package/sdk
COPY /go.mod /package/go.mod
COPY /go.sum /package/go.sum
WORKDIR /package
RUN go build -v ./...

FROM build as test
COPY /test /package/test
RUN go test -v $(go list ./... | grep -v /test/integration_tests)

FROM build as inttest
COPY /test /package/test
ARG CONDUCTOR_AUTH_KEY
ARG CONDUCTOR_AUTH_SECRET
ARG CONDUCTOR_SERVER_URL
ENV CONDUCTOR_AUTH_KEY=${CONDUCTOR_AUTH_KEY}
ENV CONDUCTOR_AUTH_SECRET=${CONDUCTOR_AUTH_SECRET}
ENV CONDUCTOR_SERVER_URL=${CONDUCTOR_SERVER_URL}
RUN go test -v ./test/integration_tests/...
