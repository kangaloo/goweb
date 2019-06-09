FROM golang:latest

RUN cd /root/ && git clone https://github.com/kangaloo/goweb.git
WORKDIR /root/goweb
RUN  cd /root/goweb && git checkout deployment && export GOPROXY=https://goproxy.io && go build

ENTRYPOINT ["/root/goweb/goweb"]
