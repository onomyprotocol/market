# Simple usage with a mounted data directory:
# > docker build -t market .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.market:/market/.market market marketd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.market:/market/.market market marketd start
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
RUN make install-standalone

# Final image
FROM alpine:edge

ENV MARKET /market

# Install ca-certificates
RUN apk add --update ca-certificates

RUN addgroup market && \
    adduser -S -G market market -h "$MARKET"

USER market

WORKDIR $MARKET

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/market-standaloned /usr/bin/marketd

# Run marketd by default, omit entrypoint to ease using container with marketcli
CMD ["marketd"]