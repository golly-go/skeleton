FROM golang:1.18.3-alpine as intialize
RUN mkdir /app


WORKDIR /app
ADD . /app

RUN mkdir /app/bin

FROM intialize as builder
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin  ./...


FROM gcr.io/distroless/base
COPY --from=builder /app/bin /app/bin
# If you have migrations or are not using gorm automigrate 
# then uncomment this out
# COPY --from=builder /app/db/migrations /app/db/migrations

WORKDIR /app

ENV PORT 3000
EXPOSE 3000

CMD ["/app/bin/golly-skeleton", "web"]
