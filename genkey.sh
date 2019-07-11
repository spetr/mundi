#!/bin/sh

openssl req \
    -new \
    -newkey rsa:4096 \
    -days 3650 \
    -nodes \
    -x509 \
    -subj "/C=CZ/ST=Czech Republic/L=Prague/O=DigitalData s.r.o./CN=*" \
    -keyout key.pem \
    -out cert.pem