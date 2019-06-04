FROM golang:latest

RUN git clone https://github.com/kangaloo/goweb.git
WORKDIR /root/goweb

RUN cd /root/goweb && go build

ENTRYPOINT ["/root/goweb/goweb"]
