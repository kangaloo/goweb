FROM golang:latest

RUN cd /root/ && git clone https://github.com/kangaloo/goweb.git
WORKDIR /root/goweb
RUN export GOPROXY=https://goproxy.io && cd /root/goweb && git checkout deployment && go build

ENTRYPOINT ["/root/goweb/goweb"]
