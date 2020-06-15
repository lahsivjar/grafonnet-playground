FROM golang:1.12.1-alpine3.9 AS bob

RUN apk add --update build-base git nodejs nodejs-npm ca-certificates && \
        rm -rf /var/cache/apk/*

WORKDIR /src

COPY . ./

RUN make clean build-release && \
        git clone https://github.com/grafana/grafonnet-lib.git dist/grafonnet-lib

FROM alpine:3.9

RUN apk add --update ca-certificates git && \
        rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=bob /src/dist /app/dist

COPY --from=bob /src/public /app/public

COPY --from=bob /src/out/grafonnet-playground /app

ENV GRAFONNET_LIB_DIR=/app/dist/grafonnet-lib

CMD ["./grafonnet-playground"]

