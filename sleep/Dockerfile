FROM alpine
RUN apk add --update bash && rm -rf /var/cache/apk/*
ADD ./sleep /sleep
EXPOSE 8080
CMD ["/sleep"]
