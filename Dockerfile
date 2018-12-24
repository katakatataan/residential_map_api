FROM golang

RUN git clone https://ff9c83fe555ed8838527699688e388dec6031bad:x-oauth-basic@github.com/katakatataan/residential_map_api.git /go/src/residential_map_api
WORKDIR /go/src/residential_map_api
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

CMD ["./main"]
