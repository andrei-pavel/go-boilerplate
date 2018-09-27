FROM golang:latest AS builder

# Copy git repository to guest.
ADD . /opt/go-boilerplate
WORKDIR /opt/go-boilerplate

# Compile.
CMD ["/bin/sh", "-c", "./scripts/test" ]
