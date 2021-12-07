FROM 172.16.5.27/golang/golang:1.15-alpine
MAINTAINER chenjia chenjia_java@163.com
ENV LANG C.UTF-8
RUN  mkdir /home
COPY my-singo /home/
COPY .env /home
RUN chmod 777 /home/my-singo

