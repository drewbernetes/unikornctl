# syntax=docker/dockerfile:1
FROM cgr.dev/chainguard/go:latest as go-builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./main.go
COPY cmd ./cmd/
COPY pkg ./pkg/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /tmp/unikornctl /app/main.go


FROM cgr.dev/chainguard/wolfi-base:latest

RUN apk update --no-cache
RUN apk add bash jq curl

RUN echo "unikornctl:x:1000:1000:Ecktl Non Root,,,:/home/unikornctl:" >> /etc/passwd
RUN mkdir -p /home/unikornctl
RUN chown unikornctl: -R /home/unikornctl

COPY --from=go-builder /tmp/unikornctl /usr/bin/

RUN chmod 755 /usr/bin/unikornctl
ENV PATH=${PATH}:/home/unikornctl/.local/bin/

USER 1000

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
