FROM alpine
RUN apk add --update bash && rm -rf /var/cache/apk/*
ADD ./deploy-v1 /deploy-v1
EXPOSE 8080
CMD ["/deploy-v1"]
