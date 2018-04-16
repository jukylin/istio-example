FROM alpine
RUN apk add --update bash curl && rm -rf /var/cache/apk/*
ADD ./index /index
EXPOSE 8000
CMD ["/index"]
