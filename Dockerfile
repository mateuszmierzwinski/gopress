FROM golang:1.19 as builder-environment
WORKDIR /app
COPY . .
RUN make in-docker

FROM scratch
MAINTAINER mateuszmierzwinski@gmail.com

COPY --from=builder-environment /app/bin/gopress-linux /gopress-linux
COPY --from=builder-environment /app/web /

EXPOSE 8080
ENTRYPOINT ["/gopress-linux"]
