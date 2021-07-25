FROM golang:1.16.5-alpine as dev

WORKDIR /work

FROM golang:1.16.5-alpine as build

COPY ./app/* /app

COPY /tests/* /tests

RUN go test ./tests -v

FROM alpine as runtime

COPY --from=build /app/app /

CMD ./app
