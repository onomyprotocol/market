# Simple usage with a mounted data directory:
# > docker build -t market .
# > docker run -it -v ~/.market:/market/.market onomy/market-dev init market --home /market/.market
# Copy genesis.json from dev/config to ~/.market/config and Dealer and Validator keys are in dev/config
# > docker run -it -v ~/.market:/market/.market onomy/market-dev keys add dealer --recover --home /market/.market
# > docker run -it -v ~/.market:/market/.market onomy/market-dev keys add validator --recover --home /market/.market
# > docker run -it -v ~/.market:/market/.market onomy/market-dev gentx validator 10000000000000000000stake --chain-id market --home /market/.market
# > docker run -it -v ~/.market:/market/.market onomy/market-dev collect-gentxs --home /market/.market
# > docker run -it -p 26656:26656 -p 26657:26657 -p 1317:1317 -p 9090:9090 -p 9091:9091 -d -v ~/.market:/market/.market onomy/market-dev start --home /market/.market
FROM golang:1.19-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/pendulum-labs/market

# Add source files
COPY . .
RUN pwd
RUN ls

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

ENV MARKET /market

# Install ca-certificates
RUN apk add --update ca-certificates

WORKDIR $MARKET

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/marketd /usr/bin/marketd

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 9091

# Run marketd by default, omit entrypoint to ease using container with marketcli
ENTRYPOINT ["marketd"]