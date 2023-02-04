FROM scratch
MAINTAINER mateuszmierzwinski@gmail.com

ADD bin/gopress-linux /gopress-linux
ADD web /

EXPOSE 8080
ENTRYPOINT ["/gopress-linux"]
