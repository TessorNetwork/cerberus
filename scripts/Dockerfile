ARG ALPINE_VERSION=3.10
ARG GOLANG_VERSION=1.19.2

FROM golang:${GOLANG_VERSION}
WORKDIR /go/src/github.com/TessorNetwork/cerberus/
COPY . .
RUN make linux

FROM alpine:${ALPINE_VERSION}
RUN apk update && apk add ca-certificates
COPY --from=0 /go/src/github.com/TessorNetwork/cerberus/build/cerberus-linux-amd64 /cerberusd
COPY --from=0 /go/src/github.com/TessorNetwork/cerberus/build/processor-linux-amd64 /processord
COPY --from=0 /go/src/github.com/TessorNetwork/cerberus/build/sync-linux-amd64 /syncd
COPY --from=0 /go/src/github.com/TessorNetwork/cerberus/build/rewards-linux-amd64 /rewardsd
COPY static /static
COPY configs/rewards.yml /configs/rewards.yml
COPY scripts/migrations /migrations
ENTRYPOINT [ "/cerberusd" ]
