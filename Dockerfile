FROM ubuntu:latest

WORKDIR /root/GoProjects/goweb/
COPY goweb /root/GoProjects/goweb/
RUN mkdir /root/GoProjects/goweb/conf && mkdir /root/GoProjects/goweb/templates
COPY conf /root/GoProjects/goweb/conf/
COPY templates /root/GoProjects/goweb/templates
ENTRYPOINT ["/root/GoProjects/goweb/goweb"]