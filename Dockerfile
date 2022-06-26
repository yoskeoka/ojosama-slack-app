# syntax = docker/dockerfile:1.3-labs

FROM --platform=linux/amd64 golang:1.18.3-alpine3.16 as builder

RUN <<EOF
    apk --update --no-cache \
        add \
        git \
        build-base
EOF

WORKDIR /app
COPY . /app

RUN make build

FROM --platform=linux/amd64 alpine:3.16

RUN addgroup --gid 1000 ojosama
RUN adduser --disabled-password --no-create-home --uid 1000 -G ojosama ojosama
USER 1000:1000
COPY --chown=ojosama --from=builder /app/bin/ojosama-slack-app /app/bin/ojosama-slack-app

CMD [ "/app/bin/ojosama-slack-app" ]
