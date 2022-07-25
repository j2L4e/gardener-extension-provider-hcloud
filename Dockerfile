############# builder
FROM eu.gcr.io/gardener-project/3rd/golang:1.18.4 AS builder

ENV BINARY_PATH=/go/bin
WORKDIR /go/src/github.com/23technologies/gardener-extension-provider-hcloud

COPY . .
RUN make build

############# base
FROM gcr.io/distroless/static-debian11:nonroot as base

WORKDIR /

############# gardener-extension-provider-hcloud
FROM base AS gardener-extension-provider-hcloud
LABEL org.opencontainers.image.source="https://github.com/23technologies/gardener-extension-provider-hcloud"

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-provider-hcloud /gardener-extension-provider-hcloud
COPY --from=builder /go/bin/gardener-extension-admission-hcloud /gardener-extension-admission-hcloud
ENTRYPOINT ["/gardener-extension-provider-hcloud"]
