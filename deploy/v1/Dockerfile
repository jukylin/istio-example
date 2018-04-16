FROM alpine
RUN apk add --update bash curl && rm -rf /var/cache/apk/*
ADD ./deploy /deploy-v1
EXPOSE 8080
CMD ["/deploy-v1"]
