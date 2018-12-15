FROM golang

WORKDIR /go/src/residential_map_api
RUN go get github.com/codegangsta/gin
RUN go get github.com/golang/dep/cmd/dep
COPY ./ .
RUN rm -rf vendor
RUN dep ensure

CMD ["go", "run", "main.go"]
