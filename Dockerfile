FROM golang:1.9.4

RUN git clone https://github.com/mmm888/go-wiki ${GOPATH}/src/github.com/mmm888/go-wiki

WORKDIR ${GOPATH}/src/github.com/mmm888/go-wiki
RUN go get -u github.com/golang/dep/cmd/dep; dep ensure; go build -o go-wiki

EXPOSE 8080:8080

VOLUME ["${GOPATH}/src/github.com/mmm888/go-wiki/wiki"]

ENTRYPOINT ["./go-wiki"]
