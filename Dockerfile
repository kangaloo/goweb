FROM golang:latest

WORKDIR /root
RUN git clone https://github.com/kangaloo/goweb.git
RUN cd /root/goweb && go build

ENTRYPOINT ["/root/goweb/goweb"]
