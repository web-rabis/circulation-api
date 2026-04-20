#
# Контейнер сборки
#
FROM golang:1.24 as builder


ENV CGO_ENABLED=0

COPY . /go/src/github.com/web-rabis/searcher
WORKDIR /go/src/github.com/web-rabis/searcher
RUN \
    version=git describe --abbrev=6 --always --tag; \
    echo "version=$version" && \
    cd cmd/apiserver && \
    go build -a -tags searcher -installsuffix searcher -ldflags "-X main.version=${version} -s -w" -o /go/bin/searcher -mod vendor

#
# Контейнер для получения актуальных SSL/TLS сертификатов
#
FROM alpine:3.16 as alpine
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S searcher && adduser -S searcher -G searcher

# копируем документацию
#RUN mkdir -p /usr/share/searcher
#COPY --from=builder /go/src/github.com/web-rabis/searcher/api /usr/share/api
#RUN chown -R searcher:searcher /usr/share/searcher

ENTRYPOINT [ "/bin/searcher" ]

#
# Контейнер рантайма
#
FROM scratch
COPY --from=builder /go/bin/searcher /bin/searcher

# копируем сертификаты из alpine
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs

## копируем документацию
#COPY --from=alpine /usr/share/searcher /usr/share/searcher

# копируем пользователя и группу из alpine
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER searcher

ENTRYPOINT ["/bin/searcher"]



