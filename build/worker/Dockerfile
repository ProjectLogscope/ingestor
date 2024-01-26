ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18

################################################################################

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build-cgo
RUN --mount=type=cache,target=/var/cache/apk \
  apk add --no-cache gcc libc-dev

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=1 go build -v -race \
  -buildvcs=false -trimpath -buildmode=exe \
  -o /bin/worker ./cmd/worker

################################################################################

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
RUN --mount=type=cache,target=/var/cache/apk \
  apk add --no-cache gcc libc-dev

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  go build -v \
  -buildvcs=false -trimpath -buildmode=exe \
  -o /bin/worker ./cmd/worker

################################################################################

FROM alpine:${ALPINE_VERSION} as final

RUN mkdir -p /service/log
RUN chmod 777 /service/log

ARG UID=10001
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  appuser
USER appuser

COPY --from=build --chmod=770 /bin/worker /service/worker

CMD ["/service/worker"]
