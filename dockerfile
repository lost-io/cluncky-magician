FROM golang:1.16.5-alpine as dev

WORKDIR /work

FROM golang:1.16.5-alpine as build

WORKDIR /app

COPY app/ /app/
RUN  go build -o app 

FROM alpine as runtime

COPY --from=build /app/app /

CMD ./app
