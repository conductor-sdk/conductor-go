FROM golang:1.17 as build
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
ARG KEY
ARG SECRET
ARG CONDUCTOR_SERVER_URL
ENV KEY=${KEY}
ENV SECRET=${SECRET}
ENV CONDUCTOR_SERVER_URL=${CONDUCTOR_SERVER_URL}
RUN go test -v ./test/integration_tests/...
