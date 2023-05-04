FROM yumafuu/mysqldef:latest as builder

FROM alpine:3.15.2
COPY --from=builder /usr/local/bin/mysqldef /usr/local/bin/mysqldef
RUN chmod +x /usr/local/bin/mysqldef

WORKDIR /migrations
COPY . /migrations

CMD ["sh", "./migrate.sh"]
