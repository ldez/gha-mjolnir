FROM golang:alpine3.8 as builder

RUN apk --update upgrade \
&& apk --no-cache --no-progress add make git gcc musl-dev \
&& rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/ldez/gha-mjolnir
COPY . .
RUN make build

FROM alpine:3.8
RUN apk update && apk add --no-cache --virtual ca-certificates
COPY --from=builder /go/src/github.com/ldez/gha-mjolnir/mjolnir /usr/bin/mjolnir

LABEL "name"="Mjolnir"
LABEL "com.github.actions.name"="Closes issues related to a merged pull request."
LABEL "com.github.actions.description"="This is an Action to close issues related to a merged pull request."
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="green"

LABEL "repository"="http://github.com/ldez/gha-mjolnir"
LABEL "homepage"="http://github.com/ldez/gha-mjolnir"
LABEL "maintainer"="ldez <ldez@users.noreply.github.com>"

ENTRYPOINT [ "/usr/bin/mjolnir" ]