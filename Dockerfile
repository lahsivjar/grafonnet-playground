FROM golang:1.12.1-alpine3.9 AS bob

RUN apk add --update build-base git nodejs nodejs-npm ca-certificates && \
        rm -rf /var/cache/apk/*

WORKDIR /src

COPY . ./
RUN npm install webpack

RUN make clean build-release && \
        git clone https://github.com/grafana/grafonnet-lib.git dist/grafonnet-lib && \
        git clone https://github.com/gojekfarm/grafonnet-bigquery-panel.git dist/grafonnet-bigquery-panel && \
        git clone https://github.com/grafana/grafonnet.git dist/grafonnet

FROM alpine:3.9

RUN apk add --update ca-certificates git && \
        rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=bob /src/dist /app/dist

COPY --from=bob /src/public /app/public

COPY --from=bob /src/out/grafonnet-playground /app

ENV GRAFONNET_LIB_DIRS='/app/dist/grafonnet-lib /app/dist/grafonnet-bigquery-panel app/dist/grafonnet'

CMD ["./grafonnet-playground"]

