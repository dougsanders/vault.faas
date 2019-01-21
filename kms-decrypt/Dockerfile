FROM alpine:3.8

RUN mkdir -p /home/app

RUN apk --no-cache add curl \
    && echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && cp /usr/bin/fwatchdog /home/app \
    && apk del curl --no-cache \
    && apk add go \
    && apk add ca-certificates \
    && apk add git \
    && apk add libc-dev \
    && apk add python \
    && apk add py-yaml

# Build Go workspace and environment
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
	&& chmod -R 777 "$GOPATH"

RUN go get -u go.mozilla.org/sops/cmd/sops

# Add non root user
RUN addgroup -S app && adduser app -S -G app

RUN chown app /home/app

COPY decrypt.sh /home/app
RUN chmod +x /home/app/decrypt.sh

WORKDIR /home/app

USER app

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs ./decrypt.sh"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1
CMD [ "fwatchdog" ]
