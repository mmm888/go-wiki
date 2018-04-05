FROM golang:1.9

RUN git config --global user.name mmm888; git config --global user.email silversoul.wt@gmail.com

RUN git clone https://github.com/mmm888/go-wiki ${GOPATH}/src/github.com/mmm888/go-wiki
WORKDIR ${GOPATH}/src/github.com/mmm888/go-wiki
RUN make build

EXPOSE 8080:8080
VOLUME ["${GOPATH}/src/github.com/mmm888/go-wiki/wiki"]

ENTRYPOINT ["./go-wiki"]
