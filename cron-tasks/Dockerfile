FROM alpine:latest
RUN apk update && apk upgrade
RUN apk --no-cache add curl
COPY . /usr/local/bin/
RUN chmod +x /usr/local/bin/*.sh
RUN echo ' *  *  *  *  * /usr/local/bin/check-violations.sh > /dev/null' >> /etc/crontabs/root
RUN echo ' */10  *  *  *  * /usr/local/bin/save-temp-list.sh' >> /etc/crontabs/root
ENTRYPOINT ["crond", "-f"]
