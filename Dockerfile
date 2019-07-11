FROM ubuntu:disco AS builder
RUN apt install golang-go
RUN go get 

FROM ubuntu:disco