FROM alpine
RUN apk add --update bash curl && rm -rf /var/cache/apk/*
ADD ./sleep /sleep
EXPOSE 8080
CMD ["/sleep"]
