#################################### BUILDER ###################################
FROM golang:alpine AS builder

# Install dependencies.
RUN apk add --update-cache bash git

# Copy git repository to guest.
ADD . /opt/go-boilerplate
WORKDIR /opt/go-boilerplate

# Compile.
RUN ./scripts/build

#################################### WORKER ####################################
FROM alpine:latest

# Copy executable from builder.
RUN mkdir -p /opt
COPY --from=builder /opt/go-boilerplate/go-boilerplate /opt/go-boilerplate/config.yaml /opt/

# Run.
WORKDIR /opt
CMD ["/bin/sh", "-c", "./go-boilerplate --config ./config.yaml"]
