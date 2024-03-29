## Stage 1 and 2 is copied from /DOCKER/Dockerfile
ARG ALIPNE_VERSION=3.17
ARG GOLANG_VERSION=1.21
#################################
# STAGE 1: install dependencies #
#################################
FROM golang:${GOLANG_VERSION}-alpine${ALIPNE_VERSION} AS base

RUN apk update && \
    apk upgrade && \
    apk --no-cache add bash git gmp-dev sudo cmake build-base libpcap-dev leveldb-dev && \
    rm -rf /var/cache/apk/*

WORKDIR /src/bls

# Install BLS library
COPY third_party ./third_party
RUN ./third_party/bls-signatures/build.sh

#####################################
# STAGE 2: install golang libraries #
#####################################
FROM base AS deps

WORKDIR /src/tenderhellar

# Fetch dependencies separately (for layer caching)
COPY go.mod go.sum ./
RUN go mod download

###########################
# STAGE 2a: INSTALL DELVE #
###########################
# This is a separate stage to "go clean" unneeded deps and optimize space
FROM deps AS delve

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go clean -cache -fuzzcache -testcache -modcache && \
    rm -fr /usr/local/go

#########################
# STAGE 3: BUILD SOURCE #
#########################
FROM deps AS e2e

WORKDIR /src/tenderhellar

# Install DLV debugger
ENV DEBUG ""
COPY --from=delve /go/bin/dlv /go/bin/dlv

LABEL maintainer="developers@hellar.io"

# Tenderhellar will be looking for the genesis file in /tenderhellar/config/genesis.json
# (unless you change `genesis_file` in config.toml). You can put your config.toml and
# private validator file into /tenderhellar/config.
#
# The /tenderhellar/data dir is used by tenderhellar to store state.
ENV TMHOME /tenderhellar

# Copy Tenderhellar source
# Avoid overwriting third-party libs
COPY . .
COPY test/e2e/docker/entrypoint* /usr/bin/

RUN rm -r /src/tenderhellar/third_party && ln -s /src/bls/third_party /src/tenderhellar/third_party && \
	cd test/e2e && make node && cp build/node /usr/bin/app && \
    go clean && go clean -cache -fuzzcache -testcache -modcache

# Set up runtime directory. We don't use a separate runtime image since we need
# e.g. leveldb and rocksdb which are already installed in the build image.
WORKDIR /tenderhellar
VOLUME /tenderhellar
ENV TMHOME=/tenderhellar
ENV GOTRACEBACK=crash

EXPOSE 26656 26657 26660 6060
ENTRYPOINT ["/usr/bin/entrypoint"]
CMD ["start"]
STOPSIGNAL SIGTERM
