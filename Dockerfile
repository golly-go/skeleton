FROM ubuntu:20.04

RUN mkdir /app
RUN mkdir /app/bin
RUN mkdir /app/db
RUN mkdir /app/db/migrations


COPY ./bin/* /app/bin/
COPY ./db/migrations/* /app/db/migrations/

RUN chmod +x /app/bin/*

WORKDIR /app

ENV PORT 3000

CMD ["/app/bin/golly-skeleton", "web"]
