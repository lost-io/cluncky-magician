FROM golang:1.16.5 as dev

WORKDIR /work

FROM golang:1.16.5 as build

COPY /app/* /app/
