FROM golang:alpine AS builder

# Install dependencies.
RUN apk add --update-cache bash git

# Copy git repository to guest.
ADD . /opt/go-boilerplate
WORKDIR /opt/go-boilerplate

# Compile.
CMD ["/bin/sh", "-c", "./scripts/test --in-docker" ]
