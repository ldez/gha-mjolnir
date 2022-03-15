FROM golang:alpine3.15 as builder

RUN apk --no-cache --no-progress add make git gcc musl-dev ca-certificates

WORKDIR /src/
COPY . .
RUN make build

FROM alpine:3.15
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/mjolnir /usr/bin/mjolnir

LABEL "name"="Mjolnir"
LABEL "com.github.actions.name"="Closes issues related to a merged pull request."
LABEL "com.github.actions.description"="This is an Action to close issues related to a merged pull request."
LABEL "com.github.actions.icon"="git-merge"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/ldez/gha-mjolnir"
LABEL "homepage"="http://github.com/ldez/gha-mjolnir"
LABEL "maintainer"="ldez <ldez@users.noreply.github.com>"

ENTRYPOINT [ "/usr/bin/mjolnir" ]
