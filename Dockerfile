FROM golang:1.10.7

ENV TZ Asia/Tokyo
COPY . /go/src/residential_map_api
WORKDIR /go/src/residential_map_api
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only
RUN go build main.go
RUN chmod +x main

EXPOSE 1323

CMD ["./main"]
